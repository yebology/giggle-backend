package ws

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chat struct {

	ClientId	primitive.ObjectID	`json:"clientId" bson:"_clientId" validate:"required"`
	Message 	string 				`json:"message" validate:"required"`

}