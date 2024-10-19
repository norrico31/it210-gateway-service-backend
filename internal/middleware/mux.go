package middleware

import (
	"log"
	"net/http"
	"time"
)

type ServeMux struct {
	mux *http.ServeMux
}

func NewServeMux() *ServeMux {
	return &ServeMux{mux: http.NewServeMux()}
}

func (c *ServeMux) HandleFunc(p string, h http.Handler, methods ...string) {
	c.mux.Handle(p, c.logRequest(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		// for _, method := range methods {
		// 	if r.Method == method {
		// 		return
		// 	}
		// }
		// http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	})))
}

func (c *ServeMux) logRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log request details
		log.Printf("%s %s %s in %s", r.Method, r.RequestURI, r.Proto, time.Since(start))

		h.ServeHTTP(w, r) // Call the next handler
	})
}

func (c *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.mux.ServeHTTP(w, r)
}
