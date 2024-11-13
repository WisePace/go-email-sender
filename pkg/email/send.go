package email

import (
	"bufio"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"sync"
)

const (
	smtpHost       = "smtp.gmail.com"
	smtpPort       = "587"
	maxConcurrency = 20
)

func SendEmail(to []string, subject, body string) error {
	sender := ""   // your gmail account
	password := "" // app-specific password for Gmail

	msg := "From: " + sender + "\n" +
		"To: " + to[0] + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	auth := smtp.PlainAuth("", sender, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, sender, to, []byte(msg))
	if err != nil {
		return err
	}

	fmt.Println("Email sent successfully:", to[0])
	return nil
}

func SendEmailsToValidRecipients(validEmails []string) error {
	messageFile, err := os.Open("letter.txt")
	if err != nil {
		return fmt.Errorf("error opening message file: %v", err)
	}
	defer messageFile.Close()

	var body string
	scanner := bufio.NewScanner(messageFile)
	for scanner.Scan() {
		body += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading message file: %v", err)
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxConcurrency)

	for _, email := range validEmails {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(recipient string) {
			defer wg.Done()
			defer func() { <-semaphore }()
			subject := "Wisespace"
			err := SendEmail([]string{recipient}, subject, body)
			if err != nil {
				log.Printf("Error sending email to %s: %v", recipient, err)
			}
		}(email)
	}

	wg.Wait()
	return nil
}
