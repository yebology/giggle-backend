package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {

	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	
}