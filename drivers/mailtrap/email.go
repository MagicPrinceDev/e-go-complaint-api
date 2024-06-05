package mailtrap

import (
	"bytes"
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

func (u *MailTrapApi) SendOTP(email, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", u.EMAIL_FROM)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Email Verification")

	// Local template
	// template, err := template.ParseFiles("./templates/otp.html")
	// Deployed template
	template, err := template.ParseFiles("goapp/templates/otp.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	data := struct {
		OTP string
	}{
		OTP: otp,
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
