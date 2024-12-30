package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/errors"
	"github.com/yebology/giggle-backend/model"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPost(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var posts []model.Post

	collection := database.GetDatabase().Collection("post")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return errors.GetError(c, fiber.StatusBadRequest, err.Error())
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &posts)
	if err != nil {
		return errors.GetError(c, fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully fetch post!",
		"data": fiber.Map{
			"posts": posts,
		},
	})

}