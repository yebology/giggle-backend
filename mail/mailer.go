package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)


func SendGreetingEmail(email string, username string) error {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		return err
	}

	html, err := template.ParseFiles("./view/greeting.html")
	if err != nil {
		fmt.Println(err)
		return err
	}

	var body bytes.Buffer
	err = html.Execute(&body, struct{
		Username string
	}{
		Username: username,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	SENDER_EMAIL := os.Getenv("SENDER_EMAIL")
	SENDER_PASSWORD := os.Getenv("SENDER_PASSWORD")

	mail := gomail.NewMessage()
	mail.SetHeader("From", SENDER_EMAIL)
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", "Welcome to Giggle! Let’s Get Started 🚀")
	mail.SetBody("text/html", body.String())

	dialer := gomail.NewDialer("smtp.gmail.com", 587, SENDER_EMAIL, SENDER_PASSWORD)
	err = dialer.DialAndSend(mail)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}