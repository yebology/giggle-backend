package helper

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
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

	err := collection.FindOne(ctx, filter).Decode(&group)
	if err != nil {
		return nil, err
	}

	receiverIds := append(group.GroupMemberIds, group.GroupOwnerId)

	return receiverIds, nil

}

func EncryptMessageWithAES256(message string) (string, error) {
	
	plainText := []byte(message)

	key := make([]byte, 32)
	_, err := rand.Reader.Read(key)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	dst := make([]byte, 0, len(plainText)+gcm.NonceSize()+gcm.Overhead())

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}

	cipherText := gcm.Seal(dst, nonce, plainText, nil)

	encryptMessage := hex.EncodeToString(cipherText)

	return encryptMessage, nil

}

func DecryptMessageWithAES256(encryptedMessage string) (string, error) {

	key := make([]byte, 32)
	_, err := rand.Reader.Read(key)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

	gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

	nonce := make([]byte, gcm.NonceSize())

	cipherText, err := hex.DecodeString(encryptedMessage)
	if err != nil {
		return "", err
	}

	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), err

}