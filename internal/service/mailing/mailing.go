package service

import (
	"fmt"
	dto "optitech/internal/dto/mailing"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func CreateMailingService(send dto.MailingReq) error {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	for _, email := range send.Emails {
		html, err := ReadMJML(email, send.PasswordMessage)
		if err != nil {
			return fmt.Errorf("could not read MJML: %w", err)
		}

		// Enviar correo
		err = SendEmail(email, send.Subject, html)
		if err != nil {
			return fmt.Errorf("could not send email: %w", err)
		}
	}

	return nil
}

func ReadMJML(email string, password string) (string, error) {
	mjmlContent, err := os.ReadFile("./internal/service/mailing/template.mjml")
	if err != nil {
		return "", fmt.Errorf("error reading MJML file: %w", err)
	}

	// Change variable
	mjmlTemplate := string(mjmlContent)
	mjmlTemplate = strings.ReplaceAll(mjmlTemplate, "{{user}}", email)
	mjmlTemplate = strings.ReplaceAll(mjmlTemplate, "{{password}}", password)

	tempMJMLFile := "temp.mjml"
	err = os.WriteFile(tempMJMLFile, []byte(mjmlTemplate), 0644)
	if err != nil {
		return "", fmt.Errorf("error writing temporary MJML file: %w", err)
	}

	// Change from mjml to html
	cmd := exec.Command("mjml", tempMJMLFile, "-o", "temp.html")
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error converting MJML to HTML: %w", err)
	}

	// Read
	htmlContent, err := os.ReadFile("temp.html")
	if err != nil {
		return "", fmt.Errorf("error reading HTML file: %w", err)
	}

	// Deleite temporal
	err = os.Remove(tempMJMLFile)
	if err != nil {
		return "", fmt.Errorf("error removing temporary MJML file: %w", err)
	}

	err = os.Remove("temp.html")
	if err != nil {
		return "", fmt.Errorf("error removing temporary HTML file: %w", err)
	}

	return string(htmlContent), nil
}

func SendEmail(email string, subject string, html string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_FROM"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", html)

	port, err := strconv.Atoi(os.Getenv("EMAIL_SMTP_PORT"))
	if err != nil {
		return fmt.Errorf("invalid SMTP port: %w", err)
	}

	d := gomail.NewDialer(os.Getenv("EMAIL_SMTP_HOST"), port, os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_SMTP_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("could not send email: %w", err)
	}

	return nil
}
