// config.go

package config

import (
	"fmt"
	"os"
	"strconv"
)

type EnvStruct struct {
	EmailFrom         string
	EmailSMTPHost     string
	EmailSMTPPort     int
	EmailSMTPPassword string
	JWTSecret         string
	JWTSecretPassword string
	WebUrl            string
}

type DigitalOceanStruct struct {
	DigitalOceanKey      string
	DigitalOceanSecret   string
	DigitalOceanEndpoint string
	DigitalOceanRegion   string
	DigitalOceanBucket   string
}

var Env *EnvStruct
var DigitalOcean *DigitalOceanStruct

func LoadConfig() error {
	portStr := os.Getenv("EMAIL_SMTP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("invalid SMTP port %q: %w", portStr, err)
	}

	Env = &EnvStruct{
		EmailFrom:         os.Getenv("EMAIL_FROM"),
		EmailSMTPHost:     os.Getenv("EMAIL_SMTP_HOST"),
		EmailSMTPPort:     port,
		EmailSMTPPassword: os.Getenv("EMAIL_SMTP_PASSWORD"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		JWTSecretPassword: os.Getenv("JWT_SECRET_PASSWORD"),
		WebUrl:            os.Getenv("WEB_URL"),
	}

	DigitalOcean = &DigitalOceanStruct{
		DigitalOceanKey:      os.Getenv("DIGITAL_OCEAN_KEY"),
		DigitalOceanSecret:   os.Getenv("DIGITAL_OCEAN_SECRET"),
		DigitalOceanEndpoint: os.Getenv("DIGITAL_OCEAN_ENDPOINT"),
		DigitalOceanRegion:   os.Getenv("DIGITAL_OCEAN_REGION"),
		DigitalOceanBucket:   os.Getenv("DIGITAL_OCEAN_BUCKET"),
	}

	return err
}
