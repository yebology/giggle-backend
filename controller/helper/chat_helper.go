package helper

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertToObjectId(id string) (primitive.ObjectID, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error converting senderId to objectId:", err)
		return primitive.NilObjectID, err
	} 

	return objectId, nil

}