package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/nyaxt/webpackage/go/signedexchange"
)

type Handler struct{}

func (Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// TODO(twifkak): Write.
	resp.Header().Set("Content-Type", "text/plain")
	_, err := resp.Write([]byte("hello world"))
	if err != nil {
		// TODO(twifkak): Log request details plus err info.
		// TODO(twifkak): Is it worth logging these? Maybe just connection drops.
		log.Println("Error serving request.")
	}
}

// Exposes an HTTP server. Don't run this on the open internet! Put it behind
// HTTPS if you wish to do that.
func main() {
	// TODO(twifkak): Make log output configurable.
	// TODO(twifkak): Replace with my own ServeMux.
	handler := Handler{}
	server := http.Server{
		Addr: ":8080",
		// Don't use DefaultServeMux, per
		// https://blog.cloudflare.com/exposing-go-on-the-internet/.
		Handler:           handler,
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
	// TCP keep-alive timeout on ListenAndServe is 3 minutes. To shorten,
	// follow the above Cloudflare blog.
	log.Fatal(server.ListenAndServe())
}
