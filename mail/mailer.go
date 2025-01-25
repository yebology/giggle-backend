package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

// SendGreetingEmail sends a greeting email to the specified recipient.
// It uses an HTML template file and replaces placeholders with the provided username.
func SendGreetingEmail(email string, username string) error {

	// Load environment variables from the .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Parse the HTML template file for the greeting email
	html, err := template.ParseFiles("./view/greeting.html")
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Replace template placeholders with dynamic data (e.g., username)
	var body bytes.Buffer
	err = html.Execute(&body, struct {
		Username string
	}{
		Username: username,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Load sender email and password from environment variables
	SENDER_EMAIL := os.Getenv("SENDER_EMAIL")
	SENDER_PASSWORD := os.Getenv("SENDER_PASSWORD")

	// Create a new email message
	mail := gomail.NewMessage()
	mail.SetHeader("From", SENDER_EMAIL)                                          // Set sender's email
	mail.SetHeader("To", email)                                                   // Set recipient's email
	mail.SetHeader("Subject", "Welcome to Giggle! Let’s Get Started 🚀")          // Set email subject
	mail.SetBody("text/html", body.String())                                      // Set email body with HTML content

	// Configure the SMTP dialer with Gmail's SMTP server and credentials
	dialer := gomail.NewDialer("smtp.gmail.com", 587, SENDER_EMAIL, SENDER_PASSWORD)

	// Send the email
	err = dialer.DialAndSend(mail)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil // Return nil if the email was sent successfully
}
