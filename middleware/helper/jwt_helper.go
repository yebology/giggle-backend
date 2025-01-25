package helper

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)
// ParseToken extracts the token from the request and parses it to validate its claims.
func ParseToken(c *fiber.Ctx) (jwt.MapClaims, error) {

	// Get the token from the Authorization header or query parameter
	token := c.Get("Authorization")
	if token == "" {
		token = c.Query("Authorization")
	}	

	// Define claims to hold the decoded data
	claims := jwt.MapClaims{}
	
	// The token is validated using the secret key.
	// jwt.ParseWithClaims parses the JWT token and validates its claims using the GetSecretKey function.
	// This function will call GetSecretKey and pass the token as an argument.
	parsedToken, err := jwt.ParseWithClaims(token, claims, GetSecretKey)
	if err != nil || !parsedToken.Valid {
		// If there’s an error parsing the token or it is invalid, return an error.
		return nil, err
	}

	// Return the claims if the token is valid.
	return claims, nil
	
}

// GetSecretKey retrieves the secret key used for validating the JWT.
// In this case, it returns a hardcoded secret key. In production, consider storing this securely.
func GetSecretKey(token *jwt.Token) (interface{}, error) {

	// Return the secret key used for signing the JWT. 
	// It's critical that the same key is used to sign and verify the token.
	// While the token parameter is passed, it is not used directly in this function,
	// but it's required by jwt.ParseWithClaims as a part of its validation process.
	return []byte("secret-key"), nil

}
