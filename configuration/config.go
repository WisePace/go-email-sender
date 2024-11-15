package configuration

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
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

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err)
	}

	viper.AutomaticEnv()

	parallelEmailEnv, err := strconv.Atoi(viper.GetString("PARALLEL_EMAIL_COUNT"))
	if err != nil || parallelEmailEnv == 0 {
		parallelEmailEnv = 20
	}

	return &Config{
		SMTPHost:           viper.GetString("SMTP_HOST"),
		SMTPPort:           viper.GetString("SMTP_PORT"),
		SMTPSender:         viper.GetString("SMTP_SENDER"),
		SMTPPassword:       viper.GetString("SMTP_PASSWORD"),
		ParallelEmailCount: parallelEmailEnv,
		EmailsList:         viper.GetString("EMAILS_LIST"),
		MessageFilePath:    viper.GetString("MESSAGE_FILE_PATH"),
		EmailSubject:       viper.GetString("EMAIL_SUBJECT"),
	}, nil
}
