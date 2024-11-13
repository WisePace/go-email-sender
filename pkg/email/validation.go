package email

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func GetValidEmails() ([]string, error) {
	file, err := os.Open("db.txt")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
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
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return validEmails, nil
}
