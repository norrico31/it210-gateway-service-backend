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
	})))
}

func (c *ServeMux) logRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the incoming request URL
		log.Printf("Request: %s %s", r.Method, r.RequestURI)

		h.ServeHTTP(w, r) // Call the next handler

		// Log how long the request took
		log.Printf("Handled: %s %s in %s", r.Method, r.RequestURI, time.Since(start))
	})
}

func (c *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.mux.ServeHTTP(w, r)
}
