package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Sms struct {
	ID        string    `json:"id"`
	To        string    `json:"to"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}

func SendConfirmationCodeEndpoint(w http.ResponseWriter, r *http.Request) {
	sms := Sms{
		ID:        "1",
		To:        "+46701234567",
		Content:   "Confirmation code: 0ldi2m",
		CreatedAt: time.Now(),
		CreatedBy: "API KEY",
	}
	json.NewEncoder(w).Encode(sms)
}

func main() {
	http.HandleFunc("/v1/sms/send-confirmation-code", SendConfirmationCodeEndpoint)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
