package main

import (
	"log"
	"pace-sender/configuration"
	"pace-sender/pkg/email"
)

func main() {
	config, err := configuration.Load()
	if err != nil {
		log.Fatal("Error loading email configuration: ", err)
	}

	validEmails, err := email.GetValidEmails()
	if err != nil {
		log.Fatal("Error fetching valid emails: ", err)
	}

	if err = email.SendEmailsToValidRecipients(validEmails, config); err != nil {
		log.Fatal("Error sending emails: ", err)
	}

}
