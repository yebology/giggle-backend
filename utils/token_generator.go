package utils

import (
	"github.com/dgrijalva/jwt-go"
	ws "github.com/yebology/giggle-backend/model/http"
)

// GenerateJWT generates a JSON Web Token (JWT) based on the provided user data.
// JWT consists of three parts: Header, Payload, and Signature.

// Header: This part contains metadata about the token, including the algorithm used to sign the token.
// Here, we are using the HS256 algorithm (HMAC with SHA256) for signing the token.
func GenerateJWT(user ws.User) (string, error) {
	
	// Create a new token with claims (claims hold the user information).
	// These claims will be placed in the Payload section of the JWT.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.Id,         // User's ID
		"email": user.Email,   // User's email
		"username": user.Username, // User's username
		"role": user.Role,         // User's role (e.g., admin or regular user)
	})

	// Signature: After the Header and Payload are prepared, the Signature part is calculated.
	// The Signature ensures the token has not been tampered with after creation.
	// We use a secret key ("secret-key") to calculate the token's signature.
	// The **secret key** is a crucial component in the signing process:
	// - It ensures that the token cannot be forged. Without the secret key, it would be impossible to generate a valid signature.
	// - If someone tries to tamper with the Header or Payload (e.g., by changing the user's role, email, etc.), the Signature will no longer match the altered data.
	// - When the server receives a JWT, it uses the same secret key to verify that the signature is correct and that the token has not been altered.
	// - The secret key should be kept private and not exposed to the public, as anyone with access to it could generate their own valid JWT.
	signedToken, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		// If there's an error while generating the signature, return the error
		return "", err
	}

	// Return the signed token
	return signedToken, err

}
