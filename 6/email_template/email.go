package main

import (
	"bytes"
	"net/smtp"
	"text/template" // Uses text templates to send plain text email
)

type EmailMessage struct {
	From, Subject, Body string
	To                  []string
}

type EmailCredentials struct {
	Username, Password, Server string
	Port                       int
}

// The email template as a string
const emailTemplate = `
From: {{.From}}
To: {{.To}}
Subject {{.Subject}}

{{.Body}}
`

var t *template.Template

func init() {
	t = template.New("email")
	t.Parse(emailTemplate)
}

func main() {
	// Populates a dataset with the email for the template and mail client
	message := &EmailMessage{
		From:    "me@example.com",
		To:      []string{"you@example.com"},
		Subject: "A test",
		Body:    "Just saying hi",
	}

	// Populates a buffer with the rendered message text from the template
	var body bytes.Buffer
	t.Execute(&body, message)

	// Sets up the SMTP email client
	authCreds := &EmailCredentials{
		Username: "myUsername",
		Password: "myPass",
		Server:   "smtp.example.com",
		Port:     25,
	}

	auth := smtp.PlainAuth("",
		authCreds.Username,
		authCreds.Password,
		authCreds.Server
	)

	// Send the email
	smtp.SendEmail(authCreds.Server+":"+strconv.Itoa(authCreds.Port),
		auth,
		message.From,
		message.To,
		// The bytes from the message buffer are passed in when the message is sent
		body.Bytes(),
	)
}
