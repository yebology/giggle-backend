package ws

import "go.mongodb.org/mongo-driver/bson/primitive"

type PersonalChat struct {

	SenderId	primitive.ObjectID		`json:"senderId" bson:"_senderId" validate:"required"`
	ReceiverId	primitive.ObjectID		`json:"receiverId" bson:"_receiverId" validate:"required"`
	Message 	string 					`json:"message" validate:"required"`

}