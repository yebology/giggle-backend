package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/controller/helper"
	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/errors"
	"github.com/yebology/giggle-backend/model"
	"go.mongodb.org/mongo-driver/bson"
)

func Register(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user model.User
	err := c.BodyParser(&user)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	collection := database.GetDatabase().Collection("user")
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	user, err = helper.GetUser(ctx, bson.M{"email": user.Email})
	if err != nil {
		return errors.GetError(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Register account successful!",
		"data": fiber.Map{
			"user": user,
		},
	})

}
