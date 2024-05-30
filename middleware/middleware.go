package middleware

import (
	"log"
	"net/http"
)

// LoggingMiddleware logs each request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("Response: %s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
	})
}
