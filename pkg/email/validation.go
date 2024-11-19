package email

import (
	"bufio"
	"fmt"
	"os"
	"pace-sender/configuration"
	"regexp"
)

func GetValidEmails(config *configuration.Config) ([]string, error) {
	if config.EmailsList == "" {
		config.EmailsList = "db.txt"
	}

	emailsFile, err := os.Open(config.EmailsList)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %v", config.EmailsList, err)
	}
	defer emailsFile.Close()

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	scanner := bufio.NewScanner(emailsFile)

	var validEmails []string
	for scanner.Scan() {
		email := scanner.Text()
		if emailRegex.MatchString(email) {
			validEmails = append(validEmails, email)
		} else {
			fmt.Println("Invalid email:", email)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", config.EmailsList, err)
	}

	return validEmails, nil
}
