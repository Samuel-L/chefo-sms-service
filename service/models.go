package service

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Sms struct {
	To        string    `json:"to"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}

func (sms *Sms) Create(client *mongo.Client) {
	collection := client.Database(os.Getenv("CHEFO_SERVICE_DATABASE")).Collection("sms")
	instance, err := collection.InsertOne(context.TODO(), sms)
	if err != nil {
		log.Println("An error occured when creating an instance of service.models.Sms: ", err)
	}
	log.Println("An instance of service.models.Sms has been created: ", instance)
}

type APIKey struct {
	Key       string    `json:"key"`
	IsActive  bool      `json:"is_active"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
