package main

import (
	"fmt"
	"net/smtp"
)

func emailSend(email string) {
	from := FromEmail
	password := EmailPassword

	// Receiver email address.
	to := []string{email}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	mainMessage := fmt.Sprintf("Welcome to master Academy, Your password  verification code is %v \n", randN)
	message := []byte(mainMessage)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")

}
