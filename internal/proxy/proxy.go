package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy(baseURL string) http.Handler {
	proxyURL, err := url.Parse(baseURL)
	if err != nil {
		log.Fatalf("Failed to parse proxy URL: %v", err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request URL and method
		log.Printf("Incoming request: %s %s", r.Method, r.RequestURI)

		// Log the full URL to which the proxy is forwarding
		fullURL := proxyURL.ResolveReference(r.URL)
		log.Printf("Forwarding to: %s", fullURL.String())

		// Forward the request to the target URL
		httputil.NewSingleHostReverseProxy(proxyURL).ServeHTTP(w, r)
	})
}
