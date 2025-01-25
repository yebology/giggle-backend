package helper

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// ParseToken is responsible for extracting and validating the JWT from the HTTP request.
// It looks for the token in either the Authorization header or the query parameter.
func ParseToken(c *fiber.Ctx) (jwt.MapClaims, error) {

	// Attempt to get the token from the "Authorization" header.
	token := c.Get("Authorization")
	if token == "" {
		// If the token is not found in the header, try getting it from the query string.
		token = c.Query("Authorization")
	}

	// Initialize an empty map to store the claims extracted from the token.
	claims := jwt.MapClaims{}
	
	// Parse the token and extract the claims, using the secret key from the GetSecretKey function.
	// The token is validated using the secret key.
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

	// Return the secret key used for signing the JWT. It's critical that the same key is used to sign and verify the token.
	return []byte("secret-key"), nil

}
