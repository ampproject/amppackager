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

/*
Transform invokes the transformer library on a given input doc to
generate and output transformed AMP HTML. This does not validate the
document.

See flag.Usage in main() for usage instructions.
*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	rpb "github.com/ampproject/amppackager/transformer/request"
	t "github.com/ampproject/amppackager/transformer"
)

var documentURLFlag = flag.String("url", "", "The URL of the document being processed, e.g. https://example.com/amphtml/article1234")
var configFlag = flag.String("config", "DEFAULT", "The configuration that determines the transformations to run. Valid values are DEFAULT, NONE, VALIDATION. See transformer.go for more info.")
var skipNewlineFlag = flag.Bool("noeol", false, "do not output the trailing newline")

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	// Custom usage message.
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUsage %s [OPTION] [FILE]\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, `
Examples:

# Execute with filename
$GOPATH/bin/transform /path/to/input.html

# Execute with pipe
cat /path/to/input.html | $GOPATH/bin/transform
`)
	}

	flag.Parse()
	var data []byte
	var err error
	switch flag.NArg() {
	case 0:
		data, err = ioutil.ReadAll(os.Stdin)
	case 1:
		data, err = ioutil.ReadFile(flag.Arg(0))
	default:
		log.Fatal("Input must be from stdin or file.")
	}
	checkErr(err)
	r := &rpb.Request{Html: string(data), DocumentUrl: *documentURLFlag}
	if *configFlag != "" {
		r.Config = rpb.Request_TransformersConfig(rpb.Request_TransformersConfig_value[*configFlag])
	}
	o, _, err := t.Process(r)
	checkErr(err)
	if *skipNewlineFlag {
		fmt.Print(o)
	} else {
		fmt.Println(o)
	}
}
