package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ReverseProxy(baseURL string) http.Handler {
	proxyURL, err := url.Parse(baseURL)
	if err != nil {
		log.Fatalf("Failed to parse proxy URL: %v", err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request URL and method
		log.Printf("Incoming request: %s %s", r.Method, r.RequestURI)

		// Define the prefix to ensure is included in the forwarded request
		prefix := "/api/v1"

		// Ensure the prefix is included in the URL path if missing
		if !strings.HasPrefix(r.URL.Path, prefix) {
			r.URL.Path = prefix + r.URL.Path
			log.Printf("Added prefix, new path: %s", r.URL.Path)
		}

		// Log the full URL to which the proxy is forwarding
		fullURL := proxyURL.ResolveReference(r.URL)
		log.Printf("Forwarding to: %s", fullURL.String())

		// Forward the request to the target URL using the reverse proxy
		httputil.NewSingleHostReverseProxy(proxyURL).ServeHTTP(w, r)
	})
}
