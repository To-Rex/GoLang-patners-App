package main

//send sms to phone number twilio api
import (
	"encoding/json"
	  "fmt"
  
	  "github.com/twilio/twilio-go"
	  twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
  )

func main() {
	// Your Account SID from twilio.com/console
	accountSid := "ACd7d3e0b9c9a9a1a1d7d3e0b9c9a9a1a1"
	// Your Auth Token from twilio.com/console
	authToken := "your_auth_token"

	// Twilio client
	client := twilio.NewClient(accountSid, authToken, nil)

	// Create message
	params := &twilioApi.MessageCreateParams{
		Body: "Hello from Go!",
		To:   "+12345678901",
		From: "+12345678901",
	}

	// Send message
	message, err := client.Messages.Create(params)
	if err != nil {
		fmt.Println(err)
	}

	// Print message
	json, _ := json.MarshalIndent(message, "", "  ")
	fmt.Println(string(json))
}


