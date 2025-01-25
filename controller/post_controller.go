package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/constant"
	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/global"
	"github.com/yebology/giggle-backend/model/http"
	"github.com/yebology/giggle-backend/output"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreatePost handles the creation of a new post
func CreatePost(c *fiber.Ctx) error {

	// Create a context with a timeout of 5 seconds to avoid blocking operations
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Parse the request body into a Post struct
	var post http.Post
	err := c.BodyParser(&post)
	if err != nil {
		// Return an error if the body cannot be parsed
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	// Validate the Post struct using the global validator
	err = global.GetValidator().Struct(post)
	if err != nil {
		// Return an error if the validation fails
		return output.GetError(c, fiber.StatusBadRequest, string(constant.ValidationError))
	}

	// Ensure post type and required talent values are valid
	if (post.PostType == "Hire" && post.RequiredTalent == 0) || (post.PostType == "Service" && post.RequiredTalent > 0) {
		// Return validation error if conditions are not met
		return output.GetError(c, fiber.StatusBadRequest, string(constant.ValidationError))
	}

	// Convert the post creator ID to an ObjectID
	objectId, err := primitive.ObjectIDFromHex(post.PostCreatorId.Hex())
	if err != nil {
		// Return error if ID conversion fails
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}
	post.PostCreatorId = objectId

	// Insert the new post into the database
	collection := database.GetDatabase().Collection("post")
	_, err = collection.InsertOne(ctx, post)
	if err != nil {
		// Return error if insertion fails
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToInsertData))
	}

	// Return success message after creating the post
	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully created a new post!",
		"data":    "",
	})
}

// UpdatePost handles updating an existing post by ID
func UpdatePost(c *fiber.Ctx) error {

	// Create a context with a timeout of 5 seconds to avoid blocking operations
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve the post ID from the URL parameters
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// Return an error if the ID is invalid
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	// Parse the request body into a Post struct
	var post http.Post
	err = c.BodyParser(&post)
	if err != nil {
		// Return an error if the body cannot be parsed
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	// Validate the Post struct using the global validator
	err = global.GetValidator().Struct(post)
	if err != nil {
		// Return an error if validation fails
		return output.GetError(c, fiber.StatusBadRequest, string(constant.ValidationError))
	}

	// Update the existing post in the database
	collection := database.GetDatabase().Collection("post")
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": post}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		// Return an error if the update fails
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToUpdateData))
	}

	// Return success message after updating the post
	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully updated a post!",
		"data":    "",
	})
}

// DeletePost handles deleting an existing post by ID
func DeletePost(c *fiber.Ctx) error {

	// Create a context with a timeout of 5 seconds to avoid blocking operations
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve the post ID from the URL parameters
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// Return an error if the ID is invalid
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	// Delete the post from the database
	collection := database.GetDatabase().Collection("post")
	filter := bson.M{"_id": objectId}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		// Return an error if the deletion fails
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToDeleteData))
	}

	// Return success message after deleting the post
	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully deleted a post!",
		"data":    "",
	})
}

// GetPosts retrieves all posts from the database
func GetPosts(c *fiber.Ctx) error {

	// Create a context with a timeout of 5 seconds to avoid blocking operations
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Define a slice to hold all posts
	var posts []http.Post

	// Retrieve all posts from the database
	collection := database.GetDatabase().Collection("post")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		// Return an error if fetching posts fails
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToRetrieveData))
	}
	defer cursor.Close(ctx)

	// Decode the cursor into the posts slice
	err = cursor.All(ctx, &posts)
	if err != nil {
		// Return an error if decoding the cursor fails
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToDecodeData))
	}

	// Return all posts in the response
	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully fetched all posts!",
		"data": fiber.Map{
			"posts": posts,
		},
	})
	
}
