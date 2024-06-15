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

func (u *MailTrapApi) SendOTP(email, otp, otp_type string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", u.EMAIL_FROM)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Email Verification")

	path := ""
	if otp_type == "forgot_password" {
		path = "./templates/forgot_password.html"
	} else {
		path = "./templates/register.html"
	}

	template, err := template.ParseFiles(path)
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
