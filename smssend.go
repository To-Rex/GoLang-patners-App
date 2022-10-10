package main

//send sms to phone number twilio api
import (
	"encoding/json"
	  "fmt"
  
	  "github.com/twilio/twilio-go"
	  twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
  )

func main() {
	msg := "Hello world"
	phone := "+998909999999"

	// Create a new Twilio client
	client := twilio.NewClient(os.Getenv("TWILIO_ACCOUNT_SID"), os.Getenv("TWILIO_AUTH_TOKEN"), nil)

	// Send a message
	message, err := client.Messages.Create(twilioApi.MessageCreateParams{
		Body: &msg,
		To:   &phone,
		From: &os.Getenv("TWILIO_PHONE_NUMBER"),
	})
	if err != nil {
		fmt.Println(err)
	}
	
}

