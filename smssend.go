package main

//send sms to phone number twilio api
import (
	"encoding/json"
	  "fmt"
  
	  "github.com/twilio/twilio-go"
	  twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
  )

  func sendsms() {
	  client := twilio.NewRestClient("ACc7e0d8c4e4b4f0e2c7f3b9d2c7b0e3b3", "e3e8d0e1c1c6e1a2b3d8b2b3c4d5e6f7")
	  params := &twilioApi.MessageCreateParams{
		  Body: "Hello from Twilio!",
		  To:   "+998909999999",
		  From: "+12058999999",
	  }
  
	  resp, err := client.Messages.Create(params)
	  if err != nil {
		  fmt.Println(err)
	  }
  
	  fmt.Println(resp.Sid)
  }

  func main() {
	  sendsms()
  }
  

