package ws

import "go.mongodb.org/mongo-driver/bson/primitive"

type Client struct {

	GroupId		primitive.ObjectID		`json:"groupId" bson:"_groupId" validate:"required"`
	UserId 		primitive.ObjectID		`json:"userId" bson:"_userId" validate:"required"`

}