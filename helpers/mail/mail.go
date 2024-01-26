package mail

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(subject, body, recipientEmail string) error {
	var (
		smtpHost     = os.Getenv("MAIL_SMTP_HOST")
		smtpPortStr  = os.Getenv("MAIL_SMTP_PORT")
		smtpUsername = os.Getenv("MAIL_SMTP_USERNAME")
		smtpPassword = os.Getenv("MAIL_SMTP_PASSWORD")
		senderEmail  = os.Getenv("MAIL_SENDER_EMAIL")
	)

	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		return fmt.Errorf("failed to convert SMTP port to integer: %s", err)
	}

	message := gomail.NewMessage()
	message.SetHeader("From", senderEmail)
	message.SetHeader("To", recipientEmail)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	dialer := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)

	if err := dialer.DialAndSend(message); err != nil {
		return fmt.Errorf("failed to send email: %s", err)
	}

	return nil
}
