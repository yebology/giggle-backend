package data

// GoogleUser represents a user structure returned by Google's API
type GoogleUser struct {

	// Email address of the Google user
	Email          string `json:"email"`   

	// Family name (last name) of the Google user
	Family_name    string `json:"family_name"`  
	
	// Given name (first name) of the Google user
	Given_name     string `json:"given_name"`    

	// Unique identifier for the Google user
	Id             string `json:"id"`            

	// Locale of the Google user (e.g., "en-US")
	Locale         string `json:"locale"`       
	
	// Full name of the Google user
	Name           string `json:"name"`         

	// URL of the Google user's profile picture
	Picture        string `json:"picture"`      

	// Whether the Google user's email is verified
	Verified_email bool   `json:"verified_email"` 

}