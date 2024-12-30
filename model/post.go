package model

import (
	"github.com/yebology/giggle-backend/model/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {

	Id 				primitive.ObjectID 	`json:"id" bson:"_id,omitempty"`
	PostCreatorId 	primitive.ObjectID 	`json:"postCreatorId" bson:"_postCreatorId"`
	PostImageHash 	string 				`json:"postImageHash"`
	PostCategory	string				`json:"postCategory"`
	PostName 		string 				`json:"postName"`
	PostDescription string 				`json:"postDescription"`
	PostPrice 		float64 			`json:"postPrice"`
	RequiredTalent 	*uint64 			`json:"requiredTalent,omitempty"`
	PostType 		constant.PostType 	`json:"postType"`
	PostStatus 		string 				`json:"postStatus"`

}