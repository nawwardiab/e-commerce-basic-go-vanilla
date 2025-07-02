package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(r *http.Request) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("%s\t%s\t%s\n", timestamp,
		r.Method, 
		r.RequestURI,
	)
}
// Handler – wraps the static handler and serves http Requests to static folder.
func Handler(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    Logger(r)
    next.ServeHTTP(w, r)
  })
}