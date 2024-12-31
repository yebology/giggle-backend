package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/constant"
	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/model"
	"github.com/yebology/giggle-backend/output"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPost(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var posts []model.Post

	collection := database.GetDatabase().Collection("post")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToRetrieveData))
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &posts)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToDecodeData))
	}

	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully fetched all post!",
		"data": fiber.Map{
			"posts": posts,
		},
	})

}