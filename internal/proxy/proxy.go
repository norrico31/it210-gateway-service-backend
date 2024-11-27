package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy(port, baseURL string) http.Handler {
	target := fmt.Sprintf("%s:%s", baseURL, port)
	proxyURL, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(proxyURL)
}
