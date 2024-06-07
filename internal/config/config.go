// config.go

package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Config struct {
	EmailFrom         string
	EmailSMTPHost     string
	EmailSMTPPort     int
	EmailSMTPPassword string
}

var instance *Config
var once sync.Once

func LoadConfig() (*Config, error) {
	var err error
	once.Do(func() {
		portStr := os.Getenv("EMAIL_SMTP_PORT")
		port, err := strconv.Atoi(portStr)
		if err != nil {
			err = fmt.Errorf("invalid SMTP port %q: %w", portStr, err)
			return
		}

		instance = &Config{
			EmailFrom:         os.Getenv("EMAIL_FROM"),
			EmailSMTPHost:     os.Getenv("EMAIL_SMTP_HOST"),
			EmailSMTPPort:     port,
			EmailSMTPPassword: os.Getenv("EMAIL_SMTP_PASSWORD"),
		}
	})

	return instance, err
}

func GetConfig() *Config {
	return instance
}
