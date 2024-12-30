package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/errors"
	"github.com/yebology/giggle-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePost(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var post model.Post
	err := c.BodyParser(&post)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	collection := database.GetDatabase().Collection("post")
	_, err = collection.InsertOne(ctx, post)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully create a new post!",
		"data": fiber.Map{
			"post": post,
		},
	})

}

func UpdatePost(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	var post model.Post
	err = c.BodyParser(&post)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	collection := database.GetDatabase().Collection("post")
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": post}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully update a post!",
		"data": fiber.Map{
			"post": post,
		},
	})

}
