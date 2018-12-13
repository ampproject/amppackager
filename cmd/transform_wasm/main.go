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

// +build go1.11

/*
Binary wrapper for the transformer lib, to be run via main.js.
*/
package main

import (
	"syscall/js"

	rpb "github.com/ampproject/amppackager/transformer/request"
	t "github.com/ampproject/amppackager/transformer"
)

func transform(args []js.Value /* url, html, cb(htmlout) */) {
	r := &rpb.Request{Html: args[1].String(), DocumentUrl: args[0].String(), Config: rpb.Request_DEFAULT}
	o, _, err := t.Process(r)
	if err != nil {
		panic(err)
	}
	args[2].Invoke(o + "\n")
}

func main() {
	cb := js.NewCallback(transform)
	defer cb.Release()
	done := make(chan struct{})
	donecb := js.NewCallback(func(args []js.Value) { done <- struct{}{} })
	js.Global().Get("begin").Invoke(cb, donecb)
	<-done
}
