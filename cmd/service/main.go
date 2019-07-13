package main

import (
	"context"
	"log"
	"net/http"

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

	api := &service.Service{DbClient: client}
	http.HandleFunc("/v1/sms/send-confirmation-code", api.SendConfirmationCodeHandler)

	log.Print("Service is running on port :8080!")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
