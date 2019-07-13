package main

import (
	"log"
	"net/http"
)

func main() {
	service := &Service{}
	service.Initialize()

	http.HandleFunc("/v1/sms/send-confirmation-code", service.SendConfirmationCodeHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
