package services

import (
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string) error {

	m := gomail.NewMessage()

	m.SetHeader("From", os.Getenv("EMAIL"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Welcome to Our Platform")
os
	m.SetBody("text/html", `
		<h2>User Created Successfully</h2>
		<p>Thank you for registering.</p>
	`)

	d := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		os.Getenv("EMAIL"),
		os.Getenv("EMAIL_PASS"),
	)

	if err := d.DialAndSend(m); err != nil {
		log.Println("Email send error:", err)
		return err
	}

	return nil
}
