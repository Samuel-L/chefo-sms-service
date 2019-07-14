package twilio

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Twilio struct {
	AccountSid string
	AuthToken  string
	URL        string
	From       string
}

func (twilio *Twilio) SendTextMessage(to string, body string) {
	log.Print("Sending text message via Twilio API")

	data := url.Values{}
	data.Set("To", to)
	data.Set("From", twilio.From)
	data.Set("Body", body)
	rb := *strings.NewReader(data.Encode())

	client := &http.Client{}

	request, err := http.NewRequest("POST", twilio.URL, &rb)
	if err != nil {
		log.Println("Received error when sending text message via Twilio API: ", err)
	}
	request.SetBasicAuth(twilio.AccountSid, twilio.AuthToken)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		log.Println("Received error when sending text message via Twilio API: ", err)
	}

	if response.StatusCode != http.StatusCreated {
		log.Println("Received error when sending text message via Twilio API: ", response.Status)
	}
}
