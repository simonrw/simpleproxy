package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
    upstreamURL := flag.String("upstream", "", "Upstream URL (e.g. http://192.168.80.2:4566)")
    bindHost := flag.String("host", "127.0.0.1:3050", "Host to bind to")
    flag.Parse()

    if *upstreamURL == "" {
        log.Fatalln("no upstream URL specified")
    }

    if *bindHost == "" {
        log.Fatalln("no bind host specified")
    }

    url, err := url.Parse(*upstreamURL)
    if err != nil {
        log.Fatalf("invalid upstream url: %v", err)
    }
    handler := httputil.NewSingleHostReverseProxy(url)
    log.Printf("Proxying requests to %v", url)
    log.Printf("Listening on %s", *bindHost)
    log.Fatal(http.ListenAndServe(*bindHost, handler))
}
