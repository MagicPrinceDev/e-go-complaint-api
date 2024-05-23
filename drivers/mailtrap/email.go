package mailtrap

import (
	"bytes"
	"path/filepath"
	"strconv"
	"text/template"

	"gopkg.in/gomail.v2"
)

type MailTrapApi struct {
	SMTP_HOST     string
	SMTP_PORT     string
	SMTP_USERNAME string
	SMTP_PASSWORD string
	EMAIL_FROM    string
}

func NewMailTrapApi(smtpHost, smtpPort, smtpUsername, smtpPassword, emailFrom string) *MailTrapApi {
	return &MailTrapApi{
		SMTP_HOST:     smtpHost,
		SMTP_PORT:     smtpPort,
		SMTP_USERNAME: smtpUsername,
		SMTP_PASSWORD: smtpPassword,
		EMAIL_FROM:    emailFrom,
	}
}

func (u *MailTrapApi) SendEmail(email string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", u.EMAIL_FROM)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Email Verification")

	verificationLink := "http://localhost:8080/api/v1/verify?email=" + email

	path := filepath.Join("templates", "otp.html")
	template, err := template.ParseFiles(path)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	data := struct {
		VerificationLink string
	}{
		VerificationLink: verificationLink,
	}

	err = template.Execute(&body, data)
	if err != nil {
		return err
	}
	m.SetBody("text/html", body.String())

	port, err := strconv.Atoi(u.SMTP_PORT)
	if err != nil {
		return err
	}
	d := gomail.NewDialer(u.SMTP_HOST, port, u.SMTP_USERNAME, u.SMTP_PASSWORD)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
