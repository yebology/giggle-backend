package helper

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
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

// Generate32BytesKey generates a 32-byte key using the SHA-256 hash function.
func Generate32BytesKey(senderObjectId primitive.ObjectID, receiverObjectId primitive.ObjectID) ([]byte) {

	// Concatenate the hexadecimal representations of the sender and receiver ObjectIDs,
	// then hash the result using SHA-256 to produce a 32-byte array.
	key := sha256.Sum256([]byte(senderObjectId.Hex() + receiverObjectId.Hex()))

	// Return the generated key as a byte slice.
	return key[:]

}

// EncryptMessageWithAES256 encrypts a given message using AES-256 in GCM mode.
func EncryptMessageWithAES256(message string, key []byte) (string, error) {
	
	// Convert the input message to a byte slice.
	plainText := []byte(message)

	// Create a new AES cipher block with the key (32-byte).
	block, err := aes.NewCipher(key)
	if err != nil {
		// It will return an error if the process fails.
		return "", err
	}
	
	// Create a GCM (Galois/Counter Mode) for encryption based on the cipher block.
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		// It will return an error if the process fails.
		return "", err
	}

	// Allocate space for the nonce (GCM requires a unique nonce for every encryption).
	nonce := make([]byte, gcm.NonceSize())

	// Fill the nonce with cryptographically secure random bytes.
	// Ensures the nonce buffer is fully populated with the required number of bytes.
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		// It will return an error if the process fails.
		return "", err
	}

	// Encrypt the plaintext using the generated nonce and GCM, and append it to the destination buffer.
	cipherText := gcm.Seal(nonce, nonce, plainText, nil)

	// Encode the ciphertext to a hex string for readability.
	encMessage := hex.EncodeToString(cipherText)

	// Return the encrypted message.
	return encMessage, nil

}

// DecryptMessageWithAES256 decrypts an AES-256 encrypted message using the provided key.
func DecryptMessageWithAES256(encMessage string, key []byte) (string, error) {

	// Decode the encrypted message to bytes.
	cipherText, err := hex.DecodeString(encMessage)
	if err != nil {
		// It will return an error if the process fails.
		return "", err
	}

	// Create a new AES cipher block with the key (32-byte).
	block, err := aes.NewCipher(key)
    if err != nil {
		// It will return an error if the process fails.
        return "", err
    }

	// Create a GCM (Galois/Counter Mode) for encryption based on the cipher block.
	gcm, err := cipher.NewGCM(block)
    if err != nil {
		// It will return an error if the process fails.
        return "", err
    }

	// Extract the nonce (initialization vector) from the start of the ciphertext.
	// Commonly, nonce size is 12 bytes.
	nonce := cipherText[:gcm.NonceSize()]

	// Update the ciphertext to exclude the nonce.
	cipherText = cipherText[gcm.NonceSize():]

	// Decrypt the ciphertext using the nonce and GCM.
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		// It will return an error if the process fails.
		return "", err
	}

	// Convert the decrypted plaintext to a string and return it.
	return string(plainText), nil

}