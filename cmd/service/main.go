package main

import (
	"log"
	"net/http"

	"github.com/samuel-l/chefo-sms-service/service"
)

func main() {
	api := &service.Service{}
	api.Initialize()

	http.HandleFunc("/v1/sms/send-confirmation-code", api.SendConfirmationCodeHandler)

	log.Print("Service is running on port :8080!")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
