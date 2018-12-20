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
	"syscall/js"

	rpb "github.com/ampproject/amppackager/transformer/request"
	t "github.com/ampproject/amppackager/transformer"
)

var maxLen = 4*1024*1024

var urlIn = make([]byte, 2000)
var htmlIn = make([]byte, maxLen)
var htmlOut = make([]byte, maxLen)

// Useful since println() doesn't go anywhere.
func consoleLog(msg string) {
	js.Global().Get("console").Call("log", msg)
}

// Transforms the doc specified in url/htmlIn into htmlOut. Args:
//   urlLen: int
//   htmlLen: int
//   done: function(htmlOutLen: int)
func transform(args []js.Value) {
	r := &rpb.Request{DocumentUrl: string(urlIn[:args[0].Int()]), Html: string(htmlIn[:args[1].Int()]), Config: rpb.Request_DEFAULT}
	o, _, err := t.Process(r)
	if err != nil {
		consoleLog(err.Error())
		o = ""  // Need to invoke the done callback with something well-defined.
	}
	if len(o) > maxLen {
		consoleLog("transformed doc too big (" + string(len(o)) + ") for url: " + string(urlIn))
		o = ""  // Need to invoke the done callback with something well-defined.
	}
	copy(htmlOut, o)
	args[2].Invoke(len(o))
}

// Takes a slice. Returns a JS function that takes a JS function cb1, calls cb1
// with a new TypedArray for the slice and a release function. The release
// function takes a JS function cb2, releases the TypedArray and itself, and
// calls cb2. Oh, the joys of continuation-passing style.
//
// The caller is responsible for releasing this function.
func typedArrayGetter(slice interface{}) js.Callback {
	return js.NewCallback(func(args []js.Value) {
		ta := js.TypedArrayOf(slice)
		var release js.Callback
		release = js.NewCallback(func(args []js.Value) {
			ta.Release()
			release.Release()
			args[0].Invoke()
		})
		args[0].Invoke(ta, release)
	})
}

func main() {
	transformCB := js.NewCallback(transform)
	defer transformCB.Release()

	done := make(chan struct{})
	doneCB := js.NewCallback(func(args []js.Value) { done <- struct{}{} })
	defer doneCB.Release()

	urlInCB := typedArrayGetter(urlIn)
	defer urlInCB.Release()

	htmlInCB := typedArrayGetter(htmlIn)
	defer htmlInCB.Release()

	htmlOutCB := typedArrayGetter(htmlOut)
	defer htmlOutCB.Release()

	js.Global().Get("begin").Invoke(transformCB, doneCB, urlInCB, htmlInCB, htmlOutCB, maxLen)
	<-done
}
