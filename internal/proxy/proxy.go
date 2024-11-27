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

	// Log the request being forwarded
	return httputil.NewSingleHostReverseProxy(proxyURL)
}
