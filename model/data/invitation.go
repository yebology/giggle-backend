package data

import "go.mongodb.org/mongo-driver/bson/primitive"

// Invitation represents the structure for inviting a member to a group
type Invitation struct {

	// MemberId is the unique identifier of the member being invited.
	MemberId 		primitive.ObjectID 		`json:"memberId" validate:"required"`

}