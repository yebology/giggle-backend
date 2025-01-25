package global

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// GetGoogleOauth initializes and returns the OAuth2 configuration for Google authentication.
// It loads the required credentials (Client ID, Client Secret, and Redirect URL) from environment variables.
func GetGoogleOauth() *oauth2.Config {

	// Load environment variables from the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load from .env: %s", err)
	}

	// Retrieve Google OAuth2 credentials from environment variables
	CLIENT_ID := os.Getenv("CLIENT_ID")
	CLIENT_SECRET := os.Getenv("CLIENT_SECRET")
	REDIRECT_URL := os.Getenv("REDIRECT_URL")

	// Create and configure the OAuth2 configuration object
	oauthConf := &oauth2.Config{
		ClientID:     CLIENT_ID,      // Google OAuth2 Client ID
		ClientSecret: CLIENT_SECRET,  // Google OAuth2 Client Secret
		RedirectURL:  REDIRECT_URL,   // Callback URL after authentication
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",   // Permission to access user's email
			"https://www.googleapis.com/auth/userinfo.profile", // Permission to access user's profile
		},
		Endpoint: google.Endpoint, // Google OAuth2 endpoint for token exchange
	}

	// Return the configured OAuth2 object
	return oauthConf
}
