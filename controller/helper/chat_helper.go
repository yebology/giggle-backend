package helper

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/model/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertToObjectId(id string) (primitive.ObjectID, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error converting id to objectId:", err)
		return primitive.NilObjectID, err
	} 

	return objectId, nil

}

func ConvertToObjectIdBoth(senderId string, receiverId string) (primitive.ObjectID, primitive.ObjectID, error) {

	senderObjectId, err := primitive.ObjectIDFromHex(senderId)
	if err != nil {
		log.Println("Error converting senderId to objectId:", err)
		return primitive.NilObjectID, primitive.NilObjectID, err
	} 

	receiverObjectId, err := primitive.ObjectIDFromHex(receiverId)
	if err != nil {
		log.Println("Error converting receiverId to objectId:", err)
		return primitive.NilObjectID, primitive.NilObjectID, err
	}

	return senderObjectId, receiverObjectId, nil

}

func CheckChatType(isGroupChatStr string) (string, error) {

	isGroupChat, err := strconv.ParseBool(isGroupChatStr)
	if err != nil {
		log.Println("Error while converting data:", err)
		return "", err
	}

	chatType := "Personal"
	if isGroupChat {
		chatType = "Group"
	}

	return chatType, nil

}

func GetGroupUsersId(filter bson.M) ([]primitive.ObjectID, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var group http.Group
	collection := database.GetDatabase().Collection("group")

	cursor, err := collection.Find(ctx, group)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &group)
	if err != nil {
		return nil, err
	}

	receiverIds := append(group.GroupMemberIds, group.GroupOwnerId)

	return receiverIds, nil

}