package mailing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"optitech/internal/config"
	dto "optitech/internal/dto/mailing"

	"gopkg.in/gomail.v2"
)

func SendPassword(send dto.PasswordMailingReq) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("could not load config: %w", err)
	}

	html, err := ReadMJML(send)
	if err != nil {
		return fmt.Errorf("could not read MJML: %w", err)
	}

	err = SendEmail(cfg, send.Email, send.Subject, html)
	if err != nil {
		return fmt.Errorf("could not send email: %w", err)
	}

	return nil
}

func ReadMJML(data interface{}) (string, error) {
	mjmlContent, err := os.ReadFile("./internal/service/mailing/templates/template-password.mjml")
	if err != nil {
		return "", fmt.Errorf("error reading MJML file: %w", err)
	}

	mjmlTemplate := string(mjmlContent)

	dataMap := make(map[string]interface{})
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error marshaling data: %w", err)
	}

	err = json.Unmarshal(dataBytes, &dataMap)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling data: %w", err)
	}

	for key, value := range dataMap {
		placeholder := fmt.Sprintf("{{%s}}", key)
		mjmlTemplate = strings.ReplaceAll(mjmlTemplate, placeholder, fmt.Sprintf("%v", value))
	}

	var mjmlBuffer bytes.Buffer
	mjmlBuffer.WriteString(mjmlTemplate)

	var htmlBuffer bytes.Buffer
	cmd := exec.Command("/app/cmd/cli/repository-cli", "convert-mjml")
	cmd.Stdin = &mjmlBuffer
	cmd.Stdout = &htmlBuffer

	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error converting MJML to HTML: %w", err)
	}

	return htmlBuffer.String(), nil
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
