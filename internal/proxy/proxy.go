package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy(baseURL string) http.Handler {
	proxyURL, _ := url.Parse(baseURL) // Only pass the base URL
	return httputil.NewSingleHostReverseProxy(proxyURL)
}
