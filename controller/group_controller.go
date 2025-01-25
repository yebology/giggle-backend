package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/constant"
	"github.com/yebology/giggle-backend/controller/helper"
	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/global"
	"github.com/yebology/giggle-backend/model/data"
	"github.com/yebology/giggle-backend/model/http"
	"github.com/yebology/giggle-backend/output"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateGroup handles the creation of a new group
func CreateGroup(c *fiber.Ctx) error {

	// Create a context with timeout to prevent blocking the request indefinitely
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Parse the incoming request body into a Group struct
	var group http.Group
	err := c.BodyParser(&group)
	if err != nil {
		// Return error if parsing fails, indicating that the data sent is not valid
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	// Validate the structure of the group object
	err = global.GetValidator().Struct(group)
	if err != nil {
		// Return error if validation fails (e.g., required fields missing or invalid format)
		return output.GetError(c, fiber.StatusBadRequest, string(constant.ValidationError))
	}

	// Ensure the group owner ID is a valid ObjectID
	_, err = primitive.ObjectIDFromHex(group.GroupOwnerId.Hex())
	if err != nil {
		// Return error if the group owner ID is not a valid ObjectID
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	// Insert the new group into the database
	collection := database.GetDatabase().Collection("group")
	_, err = collection.InsertOne(ctx, group)
	if err != nil {
		// Return error if insertion into the database fails
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToInsertData))
	}

	// Return a success response with a message
	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully created a new group!",
		"data":    "",
	})

}

// InviteMember handles the invitation of a new member to a group
func InviteMember(c *fiber.Ctx) error {

	// Create a context with timeout for database operations
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve the group ID from the URL parameters
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// Return error if the group ID is invalid
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	// Parse the invitation details from the request body
	var invitation data.Invitation
	err = c.BodyParser(&invitation)
	if err != nil {
		// Return error if the body of the invitation cannot be parsed
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	// Validate the invitation structure
	err = global.GetValidator().Struct(invitation)
	if err != nil {
		// Return error if the invitation data is invalid
		return output.GetError(c, fiber.StatusBadRequest, string(constant.ValidationError))
	}

	// Ensure the member ID in the invitation is valid
	_, err = primitive.ObjectIDFromHex(invitation.MemberId.Hex())
	if err != nil {
		// Return error if the member ID is invalid
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	// Retrieve the group from the database
	var group http.Group
	filter := bson.M{"_id": objectId}

	collection := database.GetDatabase().Collection("group")
	err = collection.FindOne(ctx, filter).Decode(&group)
	if err != nil {
		// Return error if the group cannot be found in the database
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToDecodeData))
	}

	// Add the new member to the group's member list
	groupMemberIds := append(group.GroupMemberIds, invitation.MemberId)
	update := bson.M{
		"$set": bson.M{
			"_groupMemberIds": groupMemberIds,
		},
	}

	// Update the group in the database with the new member
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		// Return error if updating the group fails
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToUpdateData))
	}

	// Return a success message
	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully invited new member!",
		"data":    "",
	})

}

// GetUserGroups retrieves all groups owned by a specific user
func GetUserGroups(c *fiber.Ctx) error {

	// Retrieve the user ID from the URL parameters
	id := c.Params("user_id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// Return error if the user ID is invalid
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	// Define a filter to find groups owned by the user
	filter := bson.M{"_groupOwnerId": objectId}

	// Retrieve groups from the database using the helper function
	groups, err := helper.GetGroupByFilter(filter)
	if err != nil {
		// Return error if the groups cannot be retrieved from the database
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToRetrieveData))
	}

	// Return the list of groups in the response
	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully fetched user groups!",
		"data": fiber.Map{
			"groups": groups,
		},
	})
	
}