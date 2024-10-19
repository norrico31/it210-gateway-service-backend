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

func (c *ServeMux) HandleFunc(p string, h http.Handler) {
	c.mux.Handle(p, c.logRequest(h))
}

func (c *ServeMux) logRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := time.Now()

		log.Printf("%s %s%s in %s", r.Method, r.Host, r.URL, time.Since(s))

		h.ServeHTTP(w, r)
	})
}

func (c *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.mux.ServeHTTP(w, r)
}
