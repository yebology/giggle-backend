package http

import (
	"github.com/yebology/giggle-backend/constant/post"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {

	Id 				primitive.ObjectID 	`json:"id" bson:"_id,omitempty"`
	PostCreatorId 	primitive.ObjectID 	`json:"postCreatorId" bson:"_postCreatorId"`
	PostImageHash 	string 				`json:"postImageHash" validate:"required"`
	PostCategory	string				`json:"postCategory" validate:"required,validPostCategory"`
	PostName 		string 				`json:"postName" validate:"required,min=8,max=255"`
	PostDescription string 				`json:"postDescription" validate:"required,min=8,max=1024"`
	PostPrice 		float64 			`json:"postPrice" validate:"required"`
	RequiredTalent 	uint64 				`json:"requiredTalent,omitempty"`
	PostType 		post.Type		 	`json:"postType" validate:"required,validPostType"`
	PostStatus 		post.Status 		`json:"postStatus" validate:"required,validPostStatus"`

}