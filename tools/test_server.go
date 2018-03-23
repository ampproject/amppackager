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
	log.Fatal(http.ListenAndServe(fmt.Sprint("localhost:", *flagPort), nil))
}
