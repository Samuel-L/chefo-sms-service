package service

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

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
		// TODO: Implement
	}

	if errors := requestBody.validate(); len(errors) > 0 {
		err := map[string]interface{}{"error": errors}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	} else {
		sms := Sms{
			ID:        "1",
			To:        requestBody.PhoneNumber,
			Content:   requestBody.ConfirmationCode,
			CreatedAt: time.Now(),
			CreatedBy: "API KEY", // TODO: Implement
		}
		sms.Create(service.DbClient)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(sms)
	}
}

func (service *Service) ValidateAPIKey(key string) bool {
	log.Print("Not implemented (API key validation)")
	return true
}
