package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Group struct {

	Id 					primitive.ObjectID 		`json:"id" bson:"_id,omitempty"`
	GroupOwnerId		primitive.ObjectID		`json:"groupOwnerId" bson:"_groupOwnerId"`
	GroupMemberIds 		*[]primitive.ObjectID 	`json:"groupMemberIds" bson:"_groupMemberIds,omitempty"`
	GroupName 			string 					`json:"groupName"`
	GroupImageHash 		string 					`json:"groupImageHash"`
	GroupDescription 	string 					`json:"groupDescription"`

}