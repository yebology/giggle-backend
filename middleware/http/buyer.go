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

// BuyerMiddleware is a middleware function that ensures the user making the request is the buyer
// associated with a proposal. If the user is not the buyer, it will return a "Permission Denied" error.
func BuyerMiddleware(c *fiber.Ctx) error {

	// Set a context with a timeout for the database query (5 seconds).
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Parse the JWT token from the request to extract user claims (e.g., user ID).
	claims, err := helper.ParseToken(c)
	if err != nil {
		// If there's an error parsing the token, return an Unauthorized error.
		return output.GetError(c, fiber.StatusUnauthorized, string(constant.InvalidTokenError))
	}

	// Retrieve the proposal ID from the URL parameters and convert it to an ObjectID.
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// If the proposal ID is invalid, return a Bad Request error.
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	// Get the "proposal" collection from the database and query the proposal using the ObjectID.
	collection := database.GetDatabase().Collection("proposal")
	filter := bson.M{"_id": objectId}

	// Initialize a variable to hold the proposal document.
	var proposal http.Proposal
	err = collection.FindOne(ctx, filter).Decode(&proposal)
	if err != nil {
		// If the proposal is not found, return a Data Unavailable error.
		return output.GetError(c, fiber.StatusBadRequest, string(constant.DataUnavailableError))
	}

	// Extract the buyer ID from the proposal and convert it to a string.
	buyerId := proposal.BuyerId.Hex()

	// Retrieve the user ID from the claims and compare it with the buyer ID.
	userId, ok := claims["id"].(string)
	if !ok || userId != buyerId {
		// If the user is not the buyer, return a "Permission Denied" error.
		return output.GetError(c, fiber.StatusForbidden, string(constant.PermissionDeniedError))
	}

	// If the user is the buyer, continue to the next middleware or handler.
	return c.Next()
	
}
