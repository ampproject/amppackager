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

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var flagPackage = flag.String("package", "", "Path to package file.")
var flagPort = flag.Int("port", 8000, "Port to serve on.")

func main() {
	flag.Parse()
	if *flagPackage == "" {
		log.Fatal("please specify --package")
	}
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		if req.URL.RequestURI() != "/" {
			http.NotFound(resp, req)
			return
		}
		resp.Header().Set("Content-Type", "text/html")
		resp.Write([]byte(`
			<link rel=prefetch href=/test.wpk>
			<a href=/test.wpk>click me!</a>
		`))
	})
	http.HandleFunc("/test.wpk", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "application/signed-exchange;v=b0")
		http.ServeFile(resp, req, *flagPackage)
	})
	log.Println("Serving on port", *flagPort)
	log.Fatal(http.ListenAndServe(fmt.Sprint("localhost:", *flagPort), nil))
}
