// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build js,wasm,!go1.12

/*
Binary wrapper for the transformer lib, to be run via main.js.
*/
package main

import (
	"bytes"
	"encoding/binary"
	"syscall/js"

	rpb "github.com/ampproject/amppackager/transformer/request"
	t "github.com/ampproject/amppackager/transformer"
	"github.com/pkg/errors"
)

// Max length of a URL. Based on the de facto limit per
// https://stackoverflow.com/a/417184.
var urlInMaxLen uint32 = 2000
// Max length of an AMP document.
var htmlInMaxLen uint32 = 4*1024*1024
// Max length of the transformed AMP. Arbitrarily larger just in case. I
// suspect the transforms may increase document size in some cases (doubtful
// ever doubling it, but I'd rather overallocate a few MBs than fail
// transformation).
var htmlOutMaxLen uint32 = htmlInMaxLen * 2

// Buffers to be used for a uint32 length prefix followed by a
// non-NUL-terminated string.
var urlIn = make([]byte, urlInMaxLen + 4)
var htmlIn = make([]byte, htmlInMaxLen + 4)
var htmlOut = make([]byte, htmlOutMaxLen + 4)


func bufferToString(buf []byte) (string, error) {
	var length uint32
	if err := binary.Read(bytes.NewReader(buf), binary.BigEndian, &length); err != nil {
		return "", errors.Wrap(err, "decoding length")
	}
	return string(buf[4:length+4]), nil
}

func errorOut(msg string) {
	js.Global().Get("console").Call("log", msg)  // Useful since println() doesn't go anywhere.
	copy(htmlOut, []byte{0, 0, 0, 0})  // Set htmlOut to empty string.
}

// Transforms the doc specified in url/htmlIn into htmlOut. Takes a nullary
// done callback as its only argument.
func transform(args []js.Value) {
	var url, html string
	var err error
	if url, err = bufferToString(urlIn); err != nil {
		errorOut(err.Error())
		return
	}
	if html, err = bufferToString(htmlIn); err != nil {
		errorOut(err.Error())
		return
	}

	r := &rpb.Request{DocumentUrl: url, Html: html, Config: rpb.Request_DEFAULT}
	out, _, err := t.Process(r)
	if err != nil {
		errorOut(err.Error())
		return
	}
	if uint64(len(out)) > uint64(htmlOutMaxLen) {
		errorOut("transformed doc too big (" + string(len(out)) + ") for url: " + string(urlIn))
		return
	}
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, uint32(len(out))); err != nil {
		errorOut(err.Error())
		return
	}
	copy(htmlOut[:4], buf.Bytes())
	copy(htmlOut[4:], out)
	args[0].Invoke()
}

// Takes a max len and a slice.
//
// Returns an object with two attributes:
//   maxLen: the max len
//   getter: JS function that takes a JS function cb1, calls cb1 with a new
//     TypedArray for the slice and a release function. The release function
//     takes a JS function cb2, releases the TypedArray and itself, and calls
//     cb2. Oh, the joys of continuation-passing style.
//
// The getter is never released. It is presumed that this object will live for
// the entire duration of the WASM process.
func typedArray(maxLen uint32, slice interface{}) js.Value {
	getter := js.NewCallback(func(args []js.Value) {
		ta := js.TypedArrayOf(slice)
		var release js.Callback
		release = js.NewCallback(func(args []js.Value) {
			ta.Release()
			release.Release()
			args[0].Invoke()
		})
		args[0].Invoke(ta, release)
	})
	return js.ValueOf(map[string]interface{}{
		"maxLen": maxLen,
		"getter": getter,
	})
}

func main() {
	transformCB := js.NewCallback(transform)
	defer transformCB.Release()

	done := make(chan struct{})
	doneCB := js.NewCallback(func(args []js.Value) { done <- struct{}{} })
	defer doneCB.Release()

	// Expose a bunch of globals for use by lib.js.
	js.Global().Set("transformCB", transformCB)
	js.Global().Set("urlIn", typedArray(urlInMaxLen, urlIn))
	js.Global().Set("htmlIn", typedArray(htmlInMaxLen, htmlIn))
	js.Global().Set("htmlOut", typedArray(htmlOutMaxLen, htmlOut))

	// Invoke the JS callback, hardcoded as {global,window}.begin, once the
	// Go is ready to receive transform requests.
	js.Global().Get("begin").Invoke(doneCB)

	// Keep the Go process running until the JS calls the done callback.
	<-done
}
