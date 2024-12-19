package main

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: %s %s", r.Method, r.URL.Path, time.Now().Format(time.RFC3339))
		next.ServeHTTP(w, r)
	})
}
