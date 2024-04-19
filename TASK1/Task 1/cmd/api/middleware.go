package main

import (
	"golang.org/x/time/rate"
	"net/http"
)

var limiter = rate.NewLimiter(2, 5)

func limitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next(w, r)
	}
}
