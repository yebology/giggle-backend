package ws

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GroupChat struct {

	Chat			Chat
	ReceiverIds 	[]primitive.ObjectID	

}