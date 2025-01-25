package http

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/constant"
	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/middleware/helper"
	"github.com/yebology/giggle-backend/model/http"
	"github.com/yebology/giggle-backend/output"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PostOwnerMiddleware is a middleware function that checks if the user making the request is the owner of the post.
// If the user is not the post owner, it will return a "Permission Denied" error.
func PostOwnerMiddleware(c *fiber.Ctx) error {

	// Set a context with a timeout for the database query (5 seconds).
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Parse the JWT token from the request to get the user claims (e.g., user ID).
	claims, err := helper.ParseToken(c)
	if err != nil {
		// If there's an error parsing the token, return an Unauthorized error.
		return output.GetError(c, fiber.StatusUnauthorized, string(constant.InvalidTokenError))
	}

	// Retrieve the post ID from the URL parameters and convert it to an ObjectID.
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// If the post ID is invalid, return a Bad Request error.
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	// Get the "post" collection from the database and query the post using the ObjectID.
	collection := database.GetDatabase().Collection("post")
	filter := bson.M{"_id": objectId}

	// Initialize a variable to hold the post document.
	var post http.Post
	err = collection.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		// If the post is not found, return a Data Unavailable error.
		return output.GetError(c, fiber.StatusBadRequest, string(constant.DataUnavailableError))
	}

	// Check if the user making the request is the creator of the post.
	postCreatorId := post.PostCreatorId.Hex()

	// Retrieve the user ID from the claims and compare it with the post creator's ID.
	userId, ok := claims["id"].(string)
	if !ok || userId != postCreatorId {
		// If the user is not the post creator, return a "Permission Denied" error.
		return output.GetError(c, fiber.StatusForbidden, string(constant.PermissionDeniedError))
	}

	// If the user is the post owner, continue to the next middleware or handler.
	return c.Next()
	
}

// GroupOwnerMiddleware is a middleware function that checks if the user making the request is the owner of the group.
// If the user is not the group owner, it will return a "Permission Denied" error.
func GroupOwnerMiddleware(c *fiber.Ctx) error {

	// Set a context with a timeout for the database query (5 seconds).
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Parse the JWT token from the request to get the user claims (e.g., user ID).
	claims, err := helper.ParseToken(c)
	if err != nil {
		// If there's an error parsing the token, return an Unauthorized error.
		return output.GetError(c, fiber.StatusUnauthorized, string(constant.InvalidTokenError))
	}

	// Retrieve the group ID from the URL parameters and convert it to an ObjectID.
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// If the group ID is invalid, return a Bad Request error.
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	// Get the "group" collection from the database and query the group using the ObjectID.
	collection := database.GetDatabase().Collection("group")
	filter := bson.M{"_id": objectId}

	// Initialize a variable to hold the group document.
	var group http.Group
	err = collection.FindOne(ctx, filter).Decode(&group)
	if err != nil {
		// If the group is not found, return a Data Unavailable error.
		return output.GetError(c, fiber.StatusBadRequest, string(constant.DataUnavailableError))
	}

	// Check if the user making the request is the owner of the group.
	groupOwnerId := group.GroupOwnerId.Hex()

	// Retrieve the user ID from the claims and compare it with the group owner's ID.
	userId, ok := claims["id"].(string)
	if userId != groupOwnerId || !ok {
		// If the user is not the group owner, return a "Permission Denied" error.
		return output.GetError(c, fiber.StatusBadRequest, string(constant.PermissionDeniedError))
	}

	// If the user is the group owner, continue to the next middleware or handler.
	return c.Next()

}
