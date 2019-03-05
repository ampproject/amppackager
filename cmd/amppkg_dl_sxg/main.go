package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/pkg/errors"
)

var flagOutSXG = flag.String("out_sxg", "test.sxg", "Path to where the signed-exchange should be saved.")
var flagOutCert = flag.String("out_cert", "test.cert", "Path to where the cert-chain+cbor should be saved.")

func getSXG(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.Header.Set("Accept", "application/signed-exchange;v=b3")
	req.Header.Set("AMP-Cache-Transform", "any")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return body, nil
}

func extractCertURL(sxg []byte) (string, error) {
	exchange, err := signedexchange.ReadExchange(bytes.NewReader(sxg))
	if err != nil {
		return "", errors.WithStack(err)
	}
	// Short of implementing a structured-headers-07 parser, we simply
	// regex for the value as returned by
	// signedexchange.Signer.signatureHeaderValue().
	re, err := regexp.Compile(`"; cert-url=(".*"); cert-sha256=\*`)
	if err != nil {
		return "", errors.WithStack(err)
	}
	matches := re.FindStringSubmatch(exchange.SignatureHeaderValue)
	if matches == nil {
		return "", errors.Errorf("no cert-url found in %s", exchange.SignatureHeaderValue)
	}
	quotedCertURL := matches[1]
	certURL, err := strconv.Unquote(quotedCertURL)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return certURL, nil
}

func getCert(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return body, nil
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprint(os.Stderr, "Usage: ", os.Args[0], " <url_of_sxg>\n\n")
		fmt.Fprint(os.Stderr, "Saves a copy of the SXG and cert-chain, to be served with amppkg_test_cache.\n\n")
		flag.Usage()
		return
	}
	sxg, err := getSXG(flag.Arg(0))
	if err != nil {
		log.Fatalf("%+v", err)
	}
	certURL, err := extractCertURL(sxg)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	err = ioutil.WriteFile(*flagOutSXG, sxg, 0644)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	cert, err := getCert(certURL)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	err = ioutil.WriteFile(*flagOutCert, cert, 0644)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	fmt.Fprintf(os.Stderr, "Saved to %s and %s.\n", *flagOutSXG, *flagOutCert)
}
