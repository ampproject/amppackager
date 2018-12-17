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

// TODO(twifkak): Improve error messages everywhere.
// TODO(twifkak): Test this.
// TODO(twifkak): Document code.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/ampproject/amppackager/packager/certcache"
	"github.com/ampproject/amppackager/packager/signer"
	"github.com/ampproject/amppackager/packager/util"
	"github.com/ampproject/amppackager/packager/validitymap"
	"github.com/ampproject/amppackager/packager/rtv"
)

var flagConfig = flag.String("config", "amppkg.toml", "Path to the config toml file.")
var flagDevelopment = flag.Bool("development", false, "True if this is a development server.")

// Prints errors returned by pkg/errors with stack traces.
func die(err interface{}) { log.Fatalf("%+v", err) }

type logIntercept struct {
	handler http.Handler
}

func (this logIntercept) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// TODO(twifkak): Adopt whatever the standard format is nowadays.
	log.Println("Serving", req.URL, "to", req.RemoteAddr)
	this.handler.ServeHTTP(resp, req)
	// TODO(twifkak): Get status code from resp. This requires making a ResponseWriter wrapper.
	// TODO(twifkak): Separate the typical weblog from the detailed error log.
}

// Exposes an HTTP server. Don't run this on the open internet, for at least two reasons:
//  - It exposes an API that allows people to sign any URL as any other URL.
//  - It is in cleartext.
func main() {
	flag.Parse()
	if *flagConfig == "" {
		die("must specify --config")
	}
	configBytes, err := ioutil.ReadFile(*flagConfig)
	if err != nil {
		die(errors.Wrapf(err, "reading config at %s", *flagConfig))
	}
	config, err := util.ReadConfig(configBytes)
	if err != nil {
		die(errors.Wrapf(err, "parsing config at %s", *flagConfig))
	}

	// TODO(twifkak): Document what cert/key storage formats this accepts.
	certPem, err := ioutil.ReadFile(config.CertFile)
	if err != nil {
		die(errors.Wrapf(err, "reading %s", config.CertFile))
	}
	keyPem, err := ioutil.ReadFile(config.KeyFile)
	if err != nil {
		die(errors.Wrapf(err, "reading %s", config.KeyFile))
	}

	certs, err := signedexchange.ParseCertificates(certPem)
	if err != nil {
		die(errors.Wrapf(err, "parsing %s", config.CertFile))
	}
	if certs == nil || len(certs) == 0 {
		die(fmt.Sprintf("no cert found in %s", config.CertFile))
	}
	if !*flagDevelopment && !util.CanSignHttpExchanges(certs[0]) {
		die("cert is missing CanSignHttpExchanges extension")
	}
	// TODO(twifkak): Verify that certs[0] covers all the signing domains in the config.

	key, err := util.ParsePrivateKey(keyPem)
	if err != nil {
		die(errors.Wrapf(err, "parsing %s", config.KeyFile))
	}
	// TODO(twifkak): Verify that key matches certs[0].

	validityMap, err := validitymap.New()
	if err != nil {
		die(errors.Wrap(err, "building validity map"))
	}

	certCache := certcache.New(certs, config.OCSPCache)
	if err = certCache.Init(nil); err != nil {
		die(errors.Wrap(err, "building cert cache"))
	}
	rtvCache, err := rtv.New()
	if err != nil {
		die(errors.Wrap(err, "initializing rtv cache"))
	}
	rtvCache.StartCron()
	defer rtvCache.StopCron()

	var overrideBaseURL *url.URL
	if *flagDevelopment {
		overrideBaseURL, err = url.Parse(fmt.Sprintf("https://localhost:%d/", config.Port))
		if err != nil {
			die(errors.Wrap(err, "parsing development base URL"))
		}
	}

	packager, err := signer.New(certs[0], key, config.URLSet, rtvCache, certCache.IsHealthy,
		overrideBaseURL, /*requireHeaders=*/!*flagDevelopment)
	if err != nil {
		die(errors.Wrap(err, "building packager"))
	}

	// TODO(twifkak): Make log output configurable.
	mux := httprouter.New()
	mux.RedirectTrailingSlash = false
	mux.RedirectFixedPath = false
	mux.GET(util.ValidityMapPath, validityMap.ServeHTTP)
	mux.GET("/priv/doc", packager.ServeHTTP)
	mux.GET("/priv/doc/*signURL", packager.ServeHTTP)
	mux.GET(path.Join(util.CertURLPrefix, ":certName"), certCache.ServeHTTP)
	addr := ""
	if config.LocalOnly {
		addr = "localhost"
	}
	addr += fmt.Sprint(":", config.Port)
	server := http.Server{
		Addr: addr,
		// Don't use DefaultServeMux, per
		// https://blog.cloudflare.com/exposing-go-on-the-internet/.
		Handler:           logIntercept{mux},
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		// If needing to stream the response, disable WriteTimeout and
		// use TimeoutHandler instead, per
		// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/.
		WriteTimeout: 60 * time.Second,
		// Needs Go 1.8.
		IdleTimeout: 120 * time.Second,
		// TODO(twifkak): Specify ErrorLog?
	}

	// TODO(twifkak): Add monitoring (e.g. per the above Cloudflare blog).

	log.Println("Serving on port", config.Port)

	// TCP keep-alive timeout on ListenAndServe is 3 minutes. To shorten,
	// follow the above Cloudflare blog.

	if *flagDevelopment {
		log.Println("WARNING: Running in development, using SXG key for TLS. This won't work in production.")
		log.Fatal(server.ListenAndServeTLS(config.CertFile, config.KeyFile))
	} else {
		log.Fatal(server.ListenAndServe())
	}
}
