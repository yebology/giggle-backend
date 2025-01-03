package mail

import (
	"html/template"

	"gopkg.in/gomail.v2"
)


func SendGreetingEmail(email string, name string) error {

	html, err := template.ParseFiles("../view/greeting.html")
	if err != nil {
		return err
	}

	err = html.Execute()

	senderEmail := "yobelnathaniel12@gmail.com"
	senderPassword := ""

	mail := gomail.NewMessage()
	mail.SetHeader("From", senderEmail)
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", "Welcome to Giggle! Let’s Get Started 🚀")

	dialer := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, senderPassword)
	err = dialer.DialAndSend(mail)
	if err != nil {
		return err
	}

	return nil

}