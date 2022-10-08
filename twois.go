package main

import (
	"fmt"
	"net/smtp"
)

func sendMailSimple() {
	auth :=smtp.PlainAuth(
		"",
		"dev.dilshodjon@gmail.com",
		"soxjmnnrefcncvix",
		"smtp.gmail.com",
	)

	msg := "subject: Hello\n\nHello, this is a test email"

	err :=smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"doimjonovasadbek1002@gmail.com",
		[]string{"doimjonovasadbek1002@gmail.com"},
		[]byte(msg),
	) 
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	sendMailSimple()
}