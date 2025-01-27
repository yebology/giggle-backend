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

// CreateProposal handles the creation of a new proposal.
// It parses the incoming request body, validates the data, and inserts it into the database.
func CreateProposal(c *fiber.Ctx) error {

	// Set a timeout context for the database operation.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var proposal http.Proposal

	// Parse the request body into the proposal struct.
	err := c.BodyParser(&proposal)
	if err != nil {
		// Return an error if the request body cannot be parsed.
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	// Validate the proposal data using a global validator instance.
	err = global.GetValidator().Struct(proposal)
	if err != nil {
		// Return a validation error if the proposal data does not meet the requirements.
		return output.GetError(c, fiber.StatusBadRequest, string(constant.ValidationError))
	}

	// Insert the validated proposal into the "proposal" collection.
	collection := database.GetDatabase().Collection("proposal")
	_, err = collection.InsertOne(ctx, proposal)
	if err != nil {
		// Return an error if the insertion fails.
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToInsertData))
	}

	// Return a success response.
	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully created a new proposal!",
		"data":    "",
	})

}

// AcceptProposal updates a proposal's status to mark it as accepted by the buyer.
func AcceptProposal(c *fiber.Ctx) error {

	// Set a timeout context for the database operation.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve the proposal ID from the URL parameters and convert it to an ObjectID.
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// Return an error if the ID is invalid.
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	// Define the filter to locate the proposal and the update to apply.
	collection := database.GetDatabase().Collection("proposal")
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"acceptByBuyer": true}}

	// Perform the update operation.
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		// Return an error if the update fails.
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToUpdateData))
	}

	// Return a success response.
	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully updated a proposal status!",
		"data":    "",
	})

}

// GetProposals retrieves proposals associated with a specific user (as buyer or creator).
func GetProposals(c *fiber.Ctx) error {

	// Set a timeout context for the database operation.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve the user ID from the URL parameters and convert it to an ObjectID.
	id := c.Params("user_id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// Return an error if the ID is invalid.
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	var proposals []http.Proposal

	// Define the filter to find proposals where the user is either the creator or the buyer.
	collection := database.GetDatabase().Collection("proposal")
	filter := bson.M{
		"$or": []bson.M{
			{"_creatorId": objectId},
			{"_buyerId": objectId},
		},
	}

	// Query the database for matching proposals.
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		// Return an error if the query fails.
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToRetrieveData))
	}
	defer cursor.Close(ctx) // Ensure the cursor is closed after use.

	// Decode the query results into the proposals slice.
	err = cursor.All(ctx, &proposals)
	if err != nil {
		// Return an error if decoding the results fails.
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToDecodeData))
	}

	// Return a success response with the retrieved proposals.
	return output.GetSuccess(c, fiber.Map{
		"message":   "Successfully fetched user's proposals!",
		"proposals": proposals,
	})

}