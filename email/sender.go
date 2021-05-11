package email

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	"github.com/sendgrid/rest"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	gomail "gopkg.in/mail.v2"
)

// SendMail  send mail using smtp server
func SendMail(email Email) error {

	host := os.Getenv("SMTP_HOST")
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	email.From = username

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", email.From)

	// Set E-Mail receivers
	m.SetHeader("To", email.To)

	// Set E-Mail subject
	m.SetHeader("Subject", email.Subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody(string(email.BodyType), email.Body)

	if email.AttachURL != nil {
		m.Attach(email.AttachURL.(string))
	}

	// Settings for SMTP server
	d := gomail.NewDialer(host, port, username, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func SendGridSenderAPI(email Email) (*rest.Response, error) {

	from := mail.NewEmail(email.FromName, email.From)
	to := mail.NewEmail("Example User", email.To)

	message := mail.NewSingleEmail(from, email.Subject, to, email.Body, "")
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		return nil, err
	}
	return response, err
}
