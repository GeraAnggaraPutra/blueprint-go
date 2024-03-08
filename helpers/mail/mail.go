package mail

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(subject, body, recipientEmail string) (err error) {
	var (
		smtpHost     = os.Getenv("MAIL_SMTP_HOST")
		smtpPortStr  = os.Getenv("MAIL_SMTP_PORT")
		smtpUsername = os.Getenv("MAIL_SMTP_USERNAME")
		smtpPassword = os.Getenv("MAIL_SMTP_PASSWORD")
		senderEmail  = os.Getenv("MAIL_SENDER_EMAIL")
	)

	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Printf("failed to convert SMTP port to integer: %s", err)
		return
	}

	message := gomail.NewMessage()
	message.SetHeader("From", senderEmail)
	message.SetHeader("To", recipientEmail)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	dialer := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)

	if err = dialer.DialAndSend(message); err != nil {
		log.Printf("failed to send email: %s", err)
		return
	}

	return
}
