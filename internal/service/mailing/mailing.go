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
	// Read .env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	html, err := ReadMJML(send.Email, send.PasswordMessage)
	if err != nil {
		return fmt.Errorf("could not read MJML: %w", err)
	}

	// Send Mail
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_FROM"))
	m.SetHeader("To", send.Email)
	m.SetHeader("Subject", send.Subject)
	m.SetBody("text/html", html)

	port, err := strconv.Atoi(os.Getenv("EMAIL_SMTP_PORT"))
	if err != nil {
		panic(err)
	}

	d := gomail.NewDialer(os.Getenv("EMAIL_SMTP_HOST"), port, os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_SMTP_PASSWORD")) // Send mail

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("could not send email: %w", err)
	}

	return nil
}

func ReadMJML(email string, password string) (string, error) {
	mjmlContent, err := os.ReadFile("template.mjml")
	if err != nil {
		return "", fmt.Errorf("error reading MJML file: %w", err)
	}

	// Change variables
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

	// Read html
	htmlContent, err := os.ReadFile("temp.html")
	if err != nil {
		return "", fmt.Errorf("error reading HTML file: %w", err)
	}

	return string(htmlContent), nil
}
