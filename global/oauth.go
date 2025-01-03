package global

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetGoogleOauth() *oauth2.Config {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load from .env: %s", err)
	}

	CLIENT_ID := os.Getenv("CLIENT_ID")
	CLIENT_SECRET := os.Getenv("CLIENT_SECRET")
	REDIRECT_URL := os.Getenv("REDIRECT_URL")

	oauthConf := &oauth2.Config{
		ClientID: CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectURL: REDIRECT_URL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return oauthConf

}