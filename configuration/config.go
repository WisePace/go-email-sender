package configuration

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

type Config struct {
	SMTPHost           string
	SMTPPort           string
	SMTPSender         string
	SMTPPassword       string
	ParallelEmailCount int
	EmailsList         string
	MessageFilePath    string
	EmailSubject       string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err)
	}

	viper.AutomaticEnv()

	parallelEmailCountString := viper.GetString("PARALLEL_EMAIL_COUNT")
	parallelEmailCount := 20

	if parallelEmailCountString != "" {
		parsedValue, err := strconv.Atoi(parallelEmailCountString)
		if err != nil {
			log.Printf("Invalid PARALLEL_EMAIL_COUNT value '%s', using default: %d", parallelEmailCountString, parallelEmailCount)
		} else {
			parallelEmailCount = parsedValue
		}
	}

	return &Config{
		SMTPHost:           viper.GetString("SMTP_HOST"),
		SMTPPort:           viper.GetString("SMTP_PORT"),
		SMTPSender:         viper.GetString("SMTP_SENDER"),
		SMTPPassword:       viper.GetString("SMTP_PASSWORD"),
		ParallelEmailCount: parallelEmailCount,
		EmailsList:         viper.GetString("EMAILS_LIST"),
		MessageFilePath:    viper.GetString("MESSAGE_FILE_PATH"),
		EmailSubject:       viper.GetString("EMAIL_SUBJECT"),
	}, nil
}
