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
// Returns htmlOutLen: int.
func transform(_ js.Value, args []js.Value) interface{} {
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
	return len(o)
}

func main() {
	transformFun := js.FuncOf(transform)
	defer transformFun.Release()

	done := make(chan struct{})
	doneFun := js.FuncOf(func(_ js.Value, _ []js.Value) interface{} { done <- struct{}{}; return nil })
	defer doneFun.Release()

	urlInTA := js.TypedArrayOf(urlIn)
	defer urlInTA.Release()

	htmlInTA := js.TypedArrayOf(htmlIn)
	defer htmlInTA.Release()

	htmlOutTA := js.TypedArrayOf(htmlOut)
	defer htmlOutTA.Release()

	js.Global().Get("begin").Invoke(transformFun, doneFun, urlInTA, htmlInTA, htmlOutTA, maxLen)
	<-done
}
