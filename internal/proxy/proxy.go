package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/norrico31/it210-gateway-service-backend/config"
)

func ReverseProxy(port string) http.Handler {
	baseURL := config.Envs.BaseURL
	target := fmt.Sprintf("%s:%s", baseURL, port)
	proxyURL, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(proxyURL)
}
