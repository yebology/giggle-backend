package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Group struct {

	Id 					primitive.ObjectID 		`json:"id" bson:"_id,omitempty"`
	GroupOwnerId		primitive.ObjectID		`json:"groupOwnerId" bson:"_groupOwnerId" validate:"required"`
	GroupMemberIds 		[]primitive.ObjectID 	`json:"groupMemberIds" bson:"_groupMemberIds"`
	GroupName 			string 					`json:"groupName" validate:"required,min=8,max=100"`
	GroupImageHash 		string 					`json:"groupImageHash" validate:"required"`
	GroupDescription 	string 					`json:"groupDescription" validate:"required,min=8,max=255"`

}