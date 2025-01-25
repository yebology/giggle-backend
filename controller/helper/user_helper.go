package helper

import (
	"context"

	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/model/http"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// CheckUser retrieves a user from the database based on the provided filter
func CheckUser(ctx context.Context, filter bson.M) (http.User, error) {

	var user http.User
	collection := database.GetDatabase().Collection("user")

	// Execute the FindOne query to fetch the user based on the filter
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		// If an error occurs (e.g., user not found), return an empty user and the error
		return http.User{}, err
	}

	// Return the retrieved user and nil error if successful
	return user, nil

}

// HashPassword generates a hashed version of the input password
func HashPassword(password string) (string, error) {

	// Use bcrypt to hash the password with a default cost factor
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// If there is an error in hashing, return an empty string and the error
		return "", err
	}

	// Return the hashed password as a string
	return string(hashedPassword), nil

}

// CheckPassword compares the given plain text password with the hashed password
func CheckPassword(hashedPassword string, password string) error {

	// Compare the provided password with the hashed password using bcrypt
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	
}
