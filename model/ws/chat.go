package ws

import (
	"github.com/yebology/giggle-backend/constant/chat"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {

	SenderId		primitive.ObjectID		`json:"senderId" bson:"_senderId" validate:"required"`
	ReceiverId		primitive.ObjectID		`json:"receiverId" bson:"_receiverId" validate:"required"`
	Message 		string 					`json:"message" validate:"required"`
	ChatTimestamp	uint64					`json:"chatTimestamp" validate:"required"`
	ChatType		chat.Chat				`json:"chatType" validate:"required,validChatType"`	

}