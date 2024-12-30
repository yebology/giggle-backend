package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/controller/helper"
	"github.com/yebology/giggle-backend/database"
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
		return output.GetError(c, fiber.StatusBadRequest ,err.Error())
	}

	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, "Error while hashing password!")
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
		return output.GetError(c, fiber.StatusBadRequest, "Email or username is already taken!")
	}

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}

	user, err = helper.GetUser(ctx, bson.M{"email": user.Email})
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, "Error while generating token access!")
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
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}

	hashedPassword, err := helper.HashPassword(login.Password)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, "Error while hashing password")
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
		return output.GetError(c, fiber.StatusBadRequest, "Invalid email or password!")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, "Error while generating token access!")
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
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}

	collection := database.GetDatabase().Collection("user")
	filter := bson.M{"email": account.Email}

	var user model.User
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, "Email hasn't registered yet!")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, "Error while generating token access!")
	}

	return output.GetSuccess(c, fiber.Map{
		"message": "Email exists!",
		"data": fiber.Map{
			"user": user,
		},
		"token": token,
	})

}