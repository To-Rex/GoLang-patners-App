package main

//send sms to phone number twilio api
import (
	"encoding/json"
	  "fmt"
  
	  "github.com/twilio/twilio-go"
	  twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
  )
  
  func main() {
	  accountSid := "ACa6ae4e7a6b6a088a86d90141bef73b59"
	  authToken := "MIAsTTrWFQSBLRlQwrJtNp0hkytJfoSzPEJ9iuCh"
  
	  client := twilio.NewRestClientWithParams(twilio.ClientParams{
		  Username: accountSid,
		  Password: authToken,
	  })
  



	  
	  params := &twilioApi.CreateMessageParams{}
	  params.SetTo("+998995340313")
	  params.SetFrom("+15017250604")
	  params.SetBody("Hello from Go!")
  
	  resp, err := client.Api.CreateMessage(params)
	  if err != nil {
		  fmt.Println(err.Error())
	  } else {
		  response, _ := json.Marshal(*resp)
		  fmt.Println("Response: " + string(response))
	  }
  }