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

// +build js,wasm,go1.12

/*
Binary wrapper for the transformer lib, to be run via main.js.
*/
package main

import (
	"bytes"
	"encoding/binary"
	"os"
	"runtime"
	"runtime/pprof"
	"syscall/js"

	rpb "github.com/ampproject/amppackager/transformer/request"
	t "github.com/ampproject/amppackager/transformer"
	"github.com/pkg/errors"
)

var num = 0
var stats runtime.MemStats

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

// A replacement for println(), which doesn't go anywhere in wasm.
func consoleLog(args... interface{}) {
	js.Global().Get("console").Call("log", args...)
}

func heapStats() {
	runtime.ReadMemStats(&stats)
	inUse := stats.HeapInuse + stats.StackInuse + stats.MSpanInuse + stats.MCacheInuse
	consoleLog("go mem: alloc =", stats.Alloc, "inUse =", inUse)
}

func heapProfile(name string) {
	f, err := os.OpenFile("wasm.go." + name + ".pprof", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0660)
	if err != nil {
		consoleLog("err:", err.Error())
		return
	}
	defer f.Close()
	pprof.WriteHeapProfile(f)
}

func errorOut(msg string) {
	consoleLog(msg)
	copy(htmlOut, []byte{0, 0, 0, 0})  // Set htmlOut to empty string.
}

// Transforms the doc specified in url/htmlIn into htmlOut. Takes a nullary
// done callback as its only argument.
func transform(this js.Value, args []js.Value) interface{} {
	num++
	if (num % 2000 == 0) {
		heapStats()
	}
	var url, html string
	var err error
	if url, err = bufferToString(urlIn); err != nil {
		errorOut(err.Error())
		return nil
	}
	if html, err = bufferToString(htmlIn); err != nil {
		errorOut(err.Error())
		return nil
	}

	r := &rpb.Request{DocumentUrl: url, Html: html, Config: rpb.Request_DEFAULT}
	out, _, err := t.Process(r)
	if err != nil {
		errorOut(err.Error())
		return nil
	}
	if uint64(len(out)) > uint64(htmlOutMaxLen) {
		errorOut("transformed doc too big (" + string(len(out)) + ") for url: " + string(urlIn))
		return nil
	}
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, uint32(len(out))); err != nil {
		errorOut(err.Error())
		return nil
	}
	copy(htmlOut[:4], buf.Bytes())
	copy(htmlOut[4:], out)
	return nil
}

// Takes a max len and a slice. Returns a JS Bytes object wrapping them.
//
// The getter is a JS function that returns an object containing two attributes:
//   ta: A new TypedArray backed by the slice.
//   release: A function that releases the TypedArray and itself.
//
// The getter is never released. It is presumed that this object will live for
// the entire duration of the WASM process.
func typedArray(maxLen uint32, slice interface{}) js.Value {
	getter := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ta := js.TypedArrayOf(slice)
		var release js.Func
		release = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			ta.Release()
			release.Release()
			return nil
		})
		return map[string]interface{}{"ta": ta, "release": release}
	})
	return js.Global().Get("GoBridge").Get("Bytes").New(getter, maxLen)
}

func main() {
	transformCB := js.FuncOf(transform)
	defer transformCB.Release()

	done := make(chan struct{})
	doneCB := js.FuncOf(func(this js.Value, args []js.Value) interface{} { done <- struct{}{}; return nil })
	defer doneCB.Release()

	// Expose a bunch of globals for use by lib.js.
	js.Global().Set("transformCB", transformCB)
	js.Global().Set("urlIn", typedArray(urlInMaxLen, urlIn))
	js.Global().Set("htmlIn", typedArray(htmlInMaxLen, htmlIn))
	js.Global().Set("htmlOut", typedArray(htmlOutMaxLen, htmlOut))

	bridge := js.Global().Get("GoBridge")
	test := bridge.Get("test").Bool()

	if test {
		heapStats()
		heapProfile("before")
	}

	// Invoke the JS callback, hardcoded as {global,window}.begin, once the
	// Go is ready to receive transform requests.
	bridge.Get("begin").Invoke(doneCB)

	// Keep the Go process running until the JS calls the done callback.
	<-done

	if test {
		heapStats()
		heapProfile("after")
	}
}
