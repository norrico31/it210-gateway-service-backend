package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy(port string) http.Handler {
	target := fmt.Sprintf("http://localhost:%s", port)
	proxyURL, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(proxyURL)
}
