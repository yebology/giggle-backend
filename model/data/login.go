package data

// Login represents the structure for user login data
type Login struct {

	// UserIdentifier is the identifier used for login. It can be either an email or a username.
	// The `validate:"required,min=8"` tag ensures that this field is required and must have at least 8 characters.
	UserIdentifier 		string 		`json:"userIdentifier" validate:"required,min=8"`	

	// Password is the user's password used for authentication.
	// The `validate:"required,min=8"` tag ensures that this field is required and must have at least 8 characters.
	Password 			string 		`json:"password" validate:"required,min=8"` 		

}