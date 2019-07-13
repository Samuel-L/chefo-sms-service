package service

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
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

	// validate phone number

	return err
}

type Service struct{}

func (service *Service) Initialize() {}

func (service *Service) SendConfirmationCodeHandler(w http.ResponseWriter, r *http.Request) {
	requestBody := &ConfirmationCodeRequest{}

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(requestBody); err != nil {
		panic(err)
	}

	w.Header().Set("Content-type", "application/json")

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
			CreatedBy: "API KEY",
		}
		sms.Create()

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(sms)
	}
}