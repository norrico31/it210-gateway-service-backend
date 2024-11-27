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
		log.Printf("Incoming request: %s %s", r.Method, r.RequestURI)

		// Strip the /api/v1 prefix if it exists
		prefix := "/api/v1"
		if strings.HasPrefix(r.URL.Path, prefix) {
			r.URL.Path = r.URL.Path[len(prefix):]
			log.Printf("Stripped prefix, new path: %s", r.URL.Path)
		}

		// Log the full URL to which the proxy is forwarding
		fullURL := proxyURL.ResolveReference(r.URL)
		log.Printf("Forwarding to: %s", fullURL.String())

		httputil.NewSingleHostReverseProxy(proxyURL).ServeHTTP(w, r)
	})
}
