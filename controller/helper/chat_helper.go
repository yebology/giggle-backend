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

// ConvertToObjectId converts a string ID to a MongoDB ObjectID
func ConvertToObjectId(id string) (primitive.ObjectID, error) {

	// Convert the string ID to a MongoDB ObjectID
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// Log the error and return a NilObjectID if conversion fails
		log.Println("Error converting id to objectId:", err)
		return primitive.NilObjectID, err
	}

	// Return the converted ObjectID and nil for error if successful
	return objectId, nil
}

// ConvertToObjectIdBoth converts two string IDs (sender and receiver) to MongoDB ObjectIDs
func ConvertToObjectIdBoth(senderId string, receiverId string) (primitive.ObjectID, primitive.ObjectID, error) {

	// Convert senderId to MongoDB ObjectID
	senderObjectId, err := primitive.ObjectIDFromHex(senderId)
	if err != nil {
		// Log the error and return NilObjectIDs for both sender and receiver if conversion fails
		log.Println("Error converting senderId to objectId:", err)
		return primitive.NilObjectID, primitive.NilObjectID, err
	}

	// Convert receiverId to MongoDB ObjectID
	receiverObjectId, err := primitive.ObjectIDFromHex(receiverId)
	if err != nil {
		// Log the error and return NilObjectIDs for both sender and receiver if conversion fails
		log.Println("Error converting receiverId to objectId:", err)
		return primitive.NilObjectID, primitive.NilObjectID, err
	}

	// Return both converted ObjectIDs and nil for error if successful
	return senderObjectId, receiverObjectId, nil
}

// CheckChatType converts a string representation of a boolean to a chat type ("Personal" or "Group")
func CheckChatType(isGroupChatStr string) (string, error) {

	// Parse the string as a boolean to determine if it's a group chat
	isGroupChat, err := strconv.ParseBool(isGroupChatStr)
	if err != nil {
		// Log the error if parsing fails and return an empty string
		log.Println("Error while converting data:", err)
		return "", err
	}

	// Return "Group" if it's a group chat, otherwise return "Personal"
	chatType := "Personal"
	if isGroupChat {
		chatType = "Group"
	}

	return chatType, nil
}

// GetGroupUsersId retrieves the list of user IDs (including the group owner) from a group
func GetGroupUsersId(filter bson.M) ([]primitive.ObjectID, error) {

	// Create a context with timeout to prevent blocking the request indefinitely
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve the group document from the database
	var group http.Group
	collection := database.GetDatabase().Collection("group")

	// Decode the group from the database based on the provided filter
	err := collection.FindOne(ctx, filter).Decode(&group)
	if err != nil {
		// Return nil if there is an error decoding the group
		return nil, err
	}

	// Combine the group member IDs and group owner ID
	receiverIds := append(group.GroupMemberIds, group.GroupOwnerId)

	// Return the list of receiver IDs (group members and owner) and nil for error if successful
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