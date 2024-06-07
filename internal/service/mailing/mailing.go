package mailing

import (
	"encoding/json"
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
	mjmlContent, err := os.ReadFile("./internal/service/mailing/templates/templatePassword.mjml")
	if err != nil {
		return "", fmt.Errorf("error reading MJML file: %w", err)
	}

	mjmlTemplate := string(mjmlContent)

	// Convertir la estructura de datos a un mapa genérico
	dataMap := make(map[string]interface{})
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error marshaling data: %w", err)
	}

	err = json.Unmarshal(dataBytes, &dataMap)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling data: %w", err)
	}

	// Reemplazar los valores en la plantilla
	for key, value := range dataMap {
		placeholder := fmt.Sprintf("{{%s}}", key)
		mjmlTemplate = strings.ReplaceAll(mjmlTemplate, placeholder, fmt.Sprintf("%v", value))
	}

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

	// Delete temp.mjml and temp.html
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
