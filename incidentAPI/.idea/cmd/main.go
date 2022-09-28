package main

import (
	"fmt"
	"net/smtp"
)

func main() {

	// Sender data.
	from := "trakkemaskintrine@gmail.com"
	password := "INPUT_PASSWORD_HERE_"

	// Receiver email address.
	to := []string{
		"cleetorres68@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

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
