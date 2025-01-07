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

func PostOwnerMiddleware(c *fiber.Ctx) error {

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

	collection := database.GetDatabase().Collection("post")
	filter := bson.M{"_id": objectId}

	var post http.Post
	err = collection.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.DataUnavailableError))
	}

	postCreatorId := post.PostCreatorId.Hex()

	userId, ok := claims["id"].(string)
	if !ok || userId != postCreatorId {
		return output.GetError(c, fiber.StatusForbidden, string(constant.PermissionDeniedError))
	}

	return c.Next()

}

func GroupOwnerMiddleware(c *fiber.Ctx) error {

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

	collection := database.GetDatabase().Collection("group")
	filter := bson.M{"_id": objectId}

	var group http.Group
	err = collection.FindOne(ctx, filter).Decode(&group)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.DataUnavailableError))
	}

	groupOwnerId := group.GroupOwnerId.Hex()

	userId, ok := claims["id"].(string)
	if userId != groupOwnerId || !ok {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.PermissionDeniedError))
	}

	return c.Next()

}