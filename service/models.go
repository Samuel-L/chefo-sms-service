package service

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Sms struct {
	ID        string    `json:"id"`
	To        string    `json:"to"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}

func (sms *Sms) Create(client *mongo.Client) {
	collection := client.Database("test").Collection("sms")
	_, err := collection.InsertOne(context.TODO(), sms)
	if err != nil {
		log.Fatal(err)
	}
	sms.SendMessage()
}

func (sms *Sms) SendMessage() {
	log.Print("Not implemented (API call to Twilio API)")
}
