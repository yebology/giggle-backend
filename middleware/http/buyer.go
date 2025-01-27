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

func BuyerMiddleware(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	claims, err := helper.ParseToken(c)
	if err != nil {
		return output.GetError(c, fiber.StatusUnauthorized, string(constant.InvalidTokenError))
	}

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	collection := database.GetDatabase().Collection("proposal")
	filter := bson.M{"_id": objectId}

	var proposal http.Proposal
	err = collection.FindOne(ctx, filter).Decode(&proposal)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.DataUnavailableError))
	}

	buyerId := proposal.BuyerId.Hex()

	userId, ok := claims["id"].(string)
	if !ok || userId != buyerId {
		return output.GetError(c, fiber.StatusForbidden, string(constant.PermissionDeniedError))
	}

	return c.Next()

}