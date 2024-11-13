package email

import (
	"bufio"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"pace-sender/configuration"
	"sync"
)

func SendEmail(to []string, subject, body, sender, password, smtpHost, smtpPort string) error {
	msg := "From: " + sender + "\n" +
		"To: " + to[0] + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	auth := smtp.PlainAuth("", sender, password, smtpHost)

	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, sender, to, []byte(msg)); err != nil {
		return err
	}
	fmt.Println("Email sent successfully:", to[0])
	return nil
}

func SendEmailsToValidRecipients(validEmails []string, config *configuration.Config) error {
	messageFile, err := os.Open(config.MessageFilePath)
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
	buffer := make(chan struct{}, config.ParallelEmailCount)

	for _, email := range validEmails {
		wg.Add(1)
		buffer <- struct{}{} // Limit concurrency

		go func(recipient string) {
			defer wg.Done()
			defer func() { <-buffer }() // Release slot for the next goroutine

			err := SendEmail([]string{recipient}, config.EmailSubject, body, config.SMTPSender, config.SMTPPassword, config.SMTPHost, config.SMTPPort)
			if err != nil {
				log.Printf("Error sending email to %s: %v", recipient, err)
			}
		}(email)
	}

	wg.Wait()
	return nil
}

func InitEmailSender() (*configuration.Config, error) {
	config, err := configuration.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("could not load config: %v", err)
	}
	return config, nil
}
