package main

import (
	"log"
	"pace-sender/pkg/email"
)

func main() {
	config, err := email.InitEmailSender()
	if err != nil {
		log.Fatal("Error loading email configuration: ", err)
	}

	validEmails, err := email.GetValidEmails()
	if err != nil {
		log.Fatal("Error fetching valid emails: ", err)
	}

	err = email.SendEmailsToValidRecipients(validEmails, config)
	if err != nil {
		log.Fatal("Error sending emails: ", err)
	}
}
