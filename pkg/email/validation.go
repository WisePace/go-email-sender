package email

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"regexp"
)

func GetValidEmails() ([]string, error) {
	emailsFilePath := viper.GetString("EMAILS_LIST")
	if emailsFilePath == "" {
		emailsFilePath = "db.txt"
	}

	file, err := os.Open(emailsFilePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %v", emailsFilePath, err)
	}
	defer file.Close()

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	var validEmails []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		email := scanner.Text()
		if emailRegex.MatchString(email) {
			validEmails = append(validEmails, email)
		} else {
			fmt.Println("Invalid email:", email)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", emailsFilePath, err)
	}

	return validEmails, nil
}
