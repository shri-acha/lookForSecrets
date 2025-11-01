package utils

import (
	"github.com/shri-acha/lookForSecrets.git/config"
	"net/smtp"
	"fmt"
)

func formatRecipients(recipients []string) string {
	result := ""
	for i, recipient := range recipients {
		if i > 0 {
			result += ", "
		}
		result += recipient
	}
	return result
}

func SendMail(config config.EmailConfig, msg config.EmailMessage) error {

	// Gmail SMTP server address
	addr := config.SMTPHost + ":" + config.SMTPPort
	
	// Authentication
	auth := smtp.PlainAuth("", config.Username, config.Password, config.SMTPHost)
	
	// Compose the email
	message := []byte(
		"From: " + msg.From + "\r\n" +
		"To: " + formatRecipients(msg.To) + "\r\n" +
		"Subject: " + msg.Subject + "\r\n" +
		"\r\n" +
		msg.Body + "\r\n")
	
	// Send the email
	err := smtp.SendMail(addr, auth, msg.From, msg.To, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	
	return nil
}
