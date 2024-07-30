package service

import "gopkg.in/gomail.v2"

func SendEmail(email, subject, body string) error {

	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	senderEmail := "momentosrede@gmail.com"
	senderPassword := "m w c y k t q r d o m w i u a i"

	recipientEmail := email

	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", recipientEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(smtpHost, smtpPort, senderEmail, senderPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
