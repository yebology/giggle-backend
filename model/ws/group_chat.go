package ws

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GroupChat represents the structure of a group chat message in the application.
type GroupChat struct {

	// Chat contains the base chat information, including sender, message, timestamp, and chat type.
	Chat Chat

	// ReceiverIds is a slice of ObjectIDs representing the users who are part of the group chat
	// and are intended to receive the message.
	ReceiverIds []primitive.ObjectID `json:"receiverIds" bson:"receiverIds"`
	
}
