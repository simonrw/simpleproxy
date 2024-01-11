package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"inet.af/tcpproxy"
)

func main() {
	upstreamURL := flag.String("upstream", "", "Upstream URL (e.g. http://192.168.80.2:4566)")
	bindHost := flag.String("host", "127.0.0.1:3050", "Host to bind to")
	mode := flag.String("mode", "tcp", "Proxy mode (http or tcp)")
	flag.Parse()

	if *upstreamURL == "" {
		log.Fatalln("no upstream URL specified")
	}

	if *bindHost == "" {
		log.Fatalln("no bind host specified")
	}

	switch *mode {
	case "http":
		httpProxy(*upstreamURL, *bindHost)
	case "tcp":
		tcpProxy(*upstreamURL, *bindHost)
	default:
		log.Fatalf("invalid mode `%s` expected `http` or `tcp`", *mode)
	}
}

func httpProxy(upstreamURL string, bindHost string) {
	if !strings.HasPrefix(upstreamURL, "http://") {
		log.Fatalln("upstream URL must start with `http://`")
	}
	url, err := url.Parse(upstreamURL)
	if err != nil {
		log.Fatalf("invalid upstream url: %v", err)
	}
	handler := httputil.NewSingleHostReverseProxy(url)
	log.Printf("Proxying requests to %v", url)
	log.Printf("Listening on %s", bindHost)
	log.Fatal(http.ListenAndServe(bindHost, handler))
}

func tcpProxy(upstreamURL string, bindHost string) {
	// TODO: validations
	if strings.HasPrefix(upstreamURL, "http://") {
		log.Fatalln("upstream URL must not start with `http://`")
	}
    var p tcpproxy.Proxy
    p.AddRoute(bindHost, tcpproxy.To(upstreamURL))
    log.Fatal(p.Run())
}
