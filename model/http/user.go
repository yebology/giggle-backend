package http

import (
	"github.com/yebology/giggle-backend/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the structure of a user in the application.
type User struct {

	// Id is the unique identifier for the user, automatically generated by MongoDB.
	Id 				primitive.ObjectID 			`json:"id" bson:"_id,omitempty"`

	// Username is the unique username of the user.
	// The `validate:"required,min=8,max=20"` tag ensures the username must be provided and its length is between 8 and 20 characters.
	Username 		string 						`json:"username" bson:"username" validate:"required,min=8,max=20"`

	// Email is the email address of the user.
	// The `validate:"required,email"` tag ensures the email must be provided and follow a valid email format.
	Email 			string 						`json:"email" bson:"email" validate:"required,email"`

	// Password is the hashed password of the user.
	// The `validate:"required,min=8"` tag ensures the password must be provided and have a minimum length of 8 characters.
	Password 		string 						`json:"password" bson:"password" validate:"required,min=8"`

	// Role defines the role of the user within the application.
	// The `validate:"required,validRole"` tag ensures the role must be provided and belong to valid predefined roles.
	Role 			constant.Role 				`json:"role" bson:"role" validate:"required,validRole"`
	
}
