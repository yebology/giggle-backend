package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/controller/helper"
	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/model"
	"github.com/yebology/giggle-backend/model/constant"
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

	var existingUser model.User
	err = collection.FindOne(ctx, filter).Decode(&existingUser)
	if err == nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.DuplicateDataError))
	}

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToInsertData))
	}

	user, err = helper.GetUser(ctx, bson.M{"email": user.Email})
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToLoadUserData))
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToGenerateTokenAccess))
	}

	return output.GetSuccess(c, fiber.Map{
		"message": "Register account successful!",
		"data": fiber.Map{
			"user": user,
		},
		"token": token,
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

	hashedPassword, err := helper.HashPassword(login.Password)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToHashPassword))
	}
	login.Password = hashedPassword

	var user model.User
	user, err = helper.GetUser(ctx, bson.M{
		"$and": []bson.M{
			{"email": login.Email},
			{"password": login.Password},
		},
	})
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidAccountError))
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToGenerateTokenAccess))
	}

	return output.GetSuccess(c, fiber.Map{
		"message": "Login successful!",
		"data": fiber.Map{
			"user": user,
		},
		"token": token,
	})

}

func CheckAccount(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var account data.Account
	err := c.BodyParser(&account)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	collection := database.GetDatabase().Collection("user")
	filter := bson.M{"email": account.Email}

	var user model.User
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.UnregisteredAccountError))
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToGenerateTokenAccess))
	}

	return output.GetSuccess(c, fiber.Map{
		"message": "Email exists!",
		"data": fiber.Map{
			"user": user,
		},
		"token": token,
	})

}