package service

import (
	"log"
	"net/http"
	"time"
)

func (service *Service) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Incoming request from: ", r.Header.Get("x-api-key"))
		startTime := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Now().Sub(startTime)
		log.Printf("Handled request from: %s in %s", r.Header.Get("x-api-key"), duration)
	})
}
