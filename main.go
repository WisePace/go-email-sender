package main

import (
	"fmt"
	"log"
	"pace-sender/pkg/email"
)

func main() {
	validEmails, err := email.GetValidEmails()
	if err != nil {
		log.Fatalf("Error validating emails: %v", err)
	}

	err = email.SendEmailsToValidRecipients(validEmails)
	if err != nil {
		log.Fatalf("Error sending emails: %v", err)
	}

	if err == nil {
		fmt.Println("Letter was sent")
	}
}
