package model

import (
	"github.com/yebology/giggle-backend/model/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {

	Id 			primitive.ObjectID 		`json:"id" bson:"_id,omitempty"`
	Username 	string 					`json:"username"`
	Email 		string 					`json:"email"`
	Password 	string 					`json:"password"`
	Role		constant.Role			`json:"role"`

}