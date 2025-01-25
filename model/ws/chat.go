package ws

import (
	"github.com/yebology/giggle-backend/constant/chat"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Chat represents the structure of a chat message in the application.
type Chat struct {

	// SenderId is the unique identifier of the user who sends the message.
	// The `validate:"required"` tag ensures that the sender's ID must be provided.
	SenderId primitive.ObjectID `json:"senderId" bson:"_senderId" validate:"required"`

	// ReceiverId is the unique identifier of the user or group who receives the message.
	// The `validate:"required"` tag ensures that the receiver's ID must be provided.
	ReceiverId primitive.ObjectID `json:"receiverId" bson:"_receiverId" validate:"required"`

	// Message contains the content of the chat message.
	// The `validate:"required"` tag ensures that the message cannot be empty.
	Message string `json:"message" validate:"required"`

	// ChatTimestamp records the time the message was sent, represented as a UNIX timestamp.
	// The `validate:"required"` tag ensures the timestamp must be provided.
	ChatTimestamp uint64 `json:"chatTimestamp" validate:"required"`

	// ChatType specifies the type of chat (e.g., personal or group).
	// The `validate:"required,validChatType"` tag ensures the chat type must be provided and match a predefined set of valid types.
	ChatType chat.Chat `json:"chatType" validate:"required,validChatType"`
	
}
