package data

import "go.mongodb.org/mongo-driver/bson/primitive"

type Invitation struct {

	MemberId 		primitive.ObjectID 		`json:"memberId" bson:"_memberId"`

}