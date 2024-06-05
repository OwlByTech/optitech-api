package mailing

import (
	"fmt"
	"optitech/internal/config"
	dto "optitech/internal/dto/mailing"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/gomail.v2"
)

func SendPassword(send dto.PasswordMailingReq) error {
	cfg, err := config.LoadConfig()
	html, err := ReadMJML(send.Email, send.Password)
	if err != nil {
		return fmt.Errorf("could not read MJML: %w", err)
	}

	err = SendEmail(cfg, send.Email, send.Subject, html)
	if err != nil {
		return fmt.Errorf("could not send email: %w", err)
	}

	return nil
}

func ReadMJML(email string, password string) (string, error) {
	mjmlContent, err := os.ReadFile("./internal/service/mailing/template.mjml")
	if err != nil {
		return "", fmt.Errorf("error reading MJML file: %w", err)
	}

	mjmlTemplate := string(mjmlContent)
	mjmlTemplate = strings.ReplaceAll(mjmlTemplate, "{{user}}", email)
	mjmlTemplate = strings.ReplaceAll(mjmlTemplate, "{{password}}", password)

	tempMJMLFile := "temp.mjml"
	err = os.WriteFile(tempMJMLFile, []byte(mjmlTemplate), 0644)
	if err != nil {
		return "", fmt.Errorf("error writing temporary MJML file: %w", err)
	}

	cmd := exec.Command("mjml", tempMJMLFile, "-o", "temp.html")
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error converting MJML to HTML: %w", err)
	}

	htmlContent, err := os.ReadFile("temp.html")
	if err != nil {
		return "", fmt.Errorf("error reading HTML file: %w", err)
	}

	return string(htmlContent), nil
}

func SendEmail(cfg *config.Config, email string, subject string, html string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.EmailFrom)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", html)

	d := gomail.NewDialer(cfg.EmailSMTPHost, cfg.EmailSMTPPort, cfg.EmailFrom, cfg.EmailSMTPPassword)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("could not send email: %w", err)
	}

	return nil
}
