package model

import (
	"github.com/yebology/giggle-backend/model/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {

	Id 			primitive.ObjectID 		`json:"id" bson:"_id,omitempty"`
	Username 	string 					`json:"username" validate:"required,min=8,max=20"`
	Email 		string 					`json:"email" validate:"required,email"`
	Password 	string 					`json:"password" validate:"required,min=8"`
	Role		constant.Role			`json:"role" validate:"required,oneof=user"`

}