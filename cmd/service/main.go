package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/samuel-l/chefo-sms-service/twilio"

	"github.com/samuel-l/chefo-sms-service/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	twilioClient := &twilio.Twilio{
		AccountSid: os.Getenv("CHEFO_TWILIO_ACCOUNT_SID"),
		AuthToken:  os.Getenv("CHEFO_TWILIO_AUTH_TOKEN"),
		From:       os.Getenv("CHEFO_TWILIO_FROM_NUMBER"),
		URL: fmt.Sprintf(
			"https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json",
			os.Getenv("CHEFO_TWILIO_ACCOUNT_SID"),
		),
	}
	api := &service.Service{
		DbClient:     client,
		TwilioClient: twilioClient,
	}
	serviceHandler := http.HandlerFunc(api.SendConfirmationCodeHandler)
	http.Handle("/v1/sms/send-confirmation-code", api.LogRequest(api.ValidateAPIKey(serviceHandler)))

	log.Print("Service is running on port :8080!")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
