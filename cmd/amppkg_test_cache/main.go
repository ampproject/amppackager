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
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/WICG/webpackage/go/signedexchange"
)

var flagSXG = flag.String("sxg", "test.sxg", "Path to signed-exchange.")
var flagCert = flag.String("cert", "test.cert", "Path to cert-chain+cbor.")
var flagPort = flag.Int("port", 8000, "Port to serve on.")

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		fmt.Fprint(os.Stderr, "Usage: ", os.Args[0], " <cert_pem> <key_pem>\n\n")
		fmt.Fprint(os.Stderr, "Serves the test SXG and cert-chain as an AMP Cache might.\n")
		fmt.Fprint(os.Stderr, "Pass it a TLS certificate pair you wish to serve with.\n\n")
		flag.Usage()
		return
	}
	certPem := flag.Arg(0)
	keyPem := flag.Arg(1)
	sxgFile, err := os.Open(*flagSXG)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer sxgFile.Close()
	exchange, err := signedexchange.ReadExchange(sxgFile)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	re, err := regexp.Compile(`"; cert-url=(".*"); cert-sha256=\*`)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	exchange.SignatureHeaderValue = re.ReplaceAllString(
		exchange.SignatureHeaderValue,
		fmt.Sprintf(`"; cert-url="https://localhost:%d/test.cert"; cert-sha256=*`, *flagPort))
	var sxg bytes.Buffer
	exchange.Write(&sxg)
	sxgReader := bytes.NewReader(sxg.Bytes())
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		if req.URL.RequestURI() != "/" {
			http.NotFound(resp, req)
			return
		}
		resp.Header().Set("Content-Type", "text/html")
		resp.Write([]byte(`
			<link rel=prefetch href=/test.sxg>
			<a href=/test.sxg>click me!</a>
		`))
	})
	http.HandleFunc("/test.cert", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "application/cert-chain+cbor")
		http.ServeFile(resp, req, *flagCert)
	})
	http.HandleFunc("/test.sxg", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "application/signed-exchange;v=b3")
		http.ServeContent(resp, req, "test.sxg", time.Time{}, sxgReader)
	})
	log.Println("Serving on port", *flagPort)
	log.Fatal(http.ListenAndServeTLS(fmt.Sprint("localhost:", *flagPort), certPem, keyPem, nil))
}
