package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/constant"
	"github.com/yebology/giggle-backend/controller/helper"
	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/global"
	"github.com/yebology/giggle-backend/model"
	"github.com/yebology/giggle-backend/model/data"
	"github.com/yebology/giggle-backend/output"
	"github.com/yebology/giggle-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func Register(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user model.User
	err := c.BodyParser(&user)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	err = global.GetValidator().Struct(user)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.ValidationError))
	}

	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToHashPassword))
	}
	user.Password = string(hashedPassword)

	collection := database.GetDatabase().Collection("user")
	filter := bson.M{
		"$or": []bson.M{
			{"email": user.Email},
			{"username": user.Username},
		},
	}

	_, err = helper.CheckUser(ctx, filter)
	if err == nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.DuplicateDataError))
	}

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToInsertData))
	}

	user, err = helper.CheckUser(ctx, filter)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToLoadUserData))
	}

	jwt, err := utils.GenerateJWT(user)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToGenerateTokenAccess))
	}

	return output.GetSuccess(c, fiber.Map{
		"message": "Register account successful!",
		"data": fiber.Map{
			"user": user,
		},
		"token": jwt,
	})

}

func Login(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var login data.Login
	err := c.BodyParser(&login)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	err = global.GetValidator().Struct(login)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.ValidationError))
	}

	filter := bson.M{
		"$or": []bson.M{
			{"email": login.UserIdentifier},
			{"username": login.UserIdentifier},
		},
	}

	user, err := helper.CheckUser(ctx, filter)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidAccountError))
	}

	err = helper.CheckPassword(user.Password, login.Password)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidAccountError))
	}

	jwt, err := utils.GenerateJWT(user)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToGenerateTokenAccess))
	}

	return output.GetSuccess(c, fiber.Map{
		"message": "Login successful!",
		"data": fiber.Map{
			"user": user,
		},
		"token": jwt,
	})

}