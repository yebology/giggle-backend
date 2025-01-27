package http

import "go.mongodb.org/mongo-driver/bson/primitive"

type Proposal struct {

	Id					primitive.ObjectID		`json:"id" bson:"_id,omitempty"`

	PostId				primitive.ObjectID		`json:"postId" validate:"required"`

	CreatorId			primitive.ObjectID		`json:"creatorId" validate:"required"`

	BuyerId				primitive.ObjectID		`json:"buyerId" validate:"required"`

	FileHash			string					`json:"fileHash" validate:"required"`
	
	FinalFee			uint64					`json:"finalFee" validate:"required"`

	DaysToComplete		uint64					`json:"daysToComplete" validate:"required"`

	AcceptByBuyer		bool					`json:"acceptByBuyer" validate:"required"`

}