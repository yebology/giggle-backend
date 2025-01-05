package ws

type PersonalChat struct {

	SenderId	string			`json:"senderId" bson:"_senderId" validate:"required"`
	ReceiverId	string			`json:"receiverId" bson:"_receiverId" validate:"required"`
	Message 	string 			`json:"message" validate:"required"`

}