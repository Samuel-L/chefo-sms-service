package service

import (
	"log"
	"time"
)

type Sms struct {
	ID        string    `json:"id"`
	To        string    `json:"to"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}

func (sms *Sms) Create() *Sms {
	log.Print("Not implemented (create SMS database instance)")
	sms.SendMessage()

	return sms
}

func (sms *Sms) SendMessage() {
	log.Print("Not implemented (API call to Twilio API)")
}
