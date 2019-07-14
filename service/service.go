package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/samuel-l/chefo-sms-service/twilio"
	"github.com/ttacon/libphonenumber"
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

	if phoneNumber, parseError := libphonenumber.Parse(request.PhoneNumber, "US"); parseError != nil {
		err.Add("phone_number", fmt.Sprintf("\"%s\" is an invalid phone number", request.PhoneNumber))
	} else {
		if phoneNumberIsValid := libphonenumber.IsValidNumber(phoneNumber); phoneNumberIsValid == false {
			err.Add("phone_number", fmt.Sprintf("\"%s\" is an invalid phone number", request.PhoneNumber))
		}
	}

	return err
}

type Service struct {
	DbClient     *mongo.Client
	APIKey       string
	TwilioClient *twilio.Twilio
}

func (service *Service) SendConfirmationCodeHandler(w http.ResponseWriter, r *http.Request) {
	requestBody := &ConfirmationCodeRequest{}

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(requestBody); err != nil {
		panic(err)
	}

	w.Header().Set("Content-type", "application/json")

	if errors := requestBody.validate(); len(errors) > 0 {
		err := map[string]interface{}{"errors": errors}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	sms := Sms{
		To:        requestBody.PhoneNumber,
		Content:   fmt.Sprintf("Your confirmation code is: %s", requestBody.ConfirmationCode),
		CreatedAt: time.Now(),
		CreatedBy: service.APIKey,
	}
	sms.Create(service.DbClient)
	service.TwilioClient.SendTextMessage(sms.To, sms.Content)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]Sms{"data": sms})
	return
}
