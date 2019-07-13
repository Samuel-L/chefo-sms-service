package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConfirmationCodeRequest struct {
	ConfirmationCode string `json:"confirmation_code"`
	PhoneNumber      string `json:"phone_number"`
}

func (request *ConfirmationCodeRequest) validate() url.Values {
	err := url.Values{}

	if request.ConfirmationCode == "" {
		err.Add("confirmation_code", "This field is required")
	}

	if request.PhoneNumber == "" {
		err.Add("phone_number", "This field is required")
	}

	// TODO: validate phone number

	return err
}

type Service struct {
	DbClient *mongo.Client
}

func (service *Service) SendConfirmationCodeHandler(w http.ResponseWriter, r *http.Request) {
	requestBody := &ConfirmationCodeRequest{}

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(requestBody); err != nil {
		panic(err)
	}

	w.Header().Set("Content-type", "application/json")

	apiKey := r.Header.Get("x-api-key")
	if !service.ValidateAPIKey(apiKey) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("API key invalid")
		return
	}

	if errors := requestBody.validate(); len(errors) > 0 {
		err := map[string]interface{}{"error": errors}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	sms := Sms{
		ID:        "1",
		To:        requestBody.PhoneNumber,
		Content:   requestBody.ConfirmationCode,
		CreatedAt: time.Now(),
		CreatedBy: apiKey, // TODO: Foreign key
	}
	sms.Create(service.DbClient)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sms)
	return
}

func (service *Service) ValidateAPIKey(key string) bool {
	var result APIKey
	filter := bson.D{{"key", key}}
	collection := service.DbClient.Database("test").Collection("api_keys")
	if err := collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		log.Printf("Validation failed: \"%s\" it not a valid API key.", key)
		return false
	}
	return true
}
