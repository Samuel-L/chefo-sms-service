package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (service *Service) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		service.APIKey = r.Header.Get("x-api-key")
		if service.APIKey == "" {
			service.APIKey = "UNKNOWN"
		}

		log.Println("Incoming request from: ", service.APIKey)
		startTime := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Now().Sub(startTime)
		log.Printf("Handled request from: %s in %s", service.APIKey, duration)
	})
}

func (service *Service) ValidateAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var result APIKey
		filter := bson.D{{"key", service.APIKey}}
		collection := service.DbClient.Database("test").Collection("api_keys")
		if err := collection.FindOne(context.TODO(), filter).Decode(&result); err != nil || !result.IsActive {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("{}")
			return
		}
		log.Printf("API key \"%s\" is valid", service.APIKey)
		next.ServeHTTP(w, r)
	})
}
