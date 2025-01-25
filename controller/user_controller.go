package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/constant"
	"github.com/yebology/giggle-backend/controller/helper"
	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/global"
	"github.com/yebology/giggle-backend/mail"
	"github.com/yebology/giggle-backend/model/data"
	"github.com/yebology/giggle-backend/model/http"
	"github.com/yebology/giggle-backend/output"
	"github.com/yebology/giggle-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// Register handles user registration
func Register(c *fiber.Ctx) error {

	// Create a context with timeout to prevent blocking the request indefinitely
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Parse the incoming request body into a User struct
	var user http.User
	err := c.BodyParser(&user)
	if err != nil {
		// Return error if the data cannot be parsed correctly
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	// Validate the structure of the user object
	err = global.GetValidator().Struct(user)
	if err != nil {
		// Return error if the user data is invalid (e.g., required fields missing)
		return output.GetError(c, fiber.StatusBadRequest, string(constant.ValidationError))
	}

	// Hash the password using the helper function
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		// Return error if hashing the password fails
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToHashPassword))
	}
	user.Password = string(hashedPassword)

	// Check if the email or username already exists in the database
	collection := database.GetDatabase().Collection("user")
	filter := bson.M{
		"$or": []bson.M{
			{"email": user.Email},
			{"username": user.Username},
		},
	}

	// Check user existence
	_, err = helper.CheckUser(ctx, filter)
	if err == nil {
		// Return error if the email or username already exists
		return output.GetError(c, fiber.StatusBadRequest, string(constant.DuplicateDataError))
	}

	// Insert the new user into the database
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		// Return error if the insertion into the database fails
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToInsertData))
	}

	// Retrieve the newly created user from the database
	user, err = helper.CheckUser(ctx, filter)
	if err != nil {
		// Return error if loading the user data fails
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToLoadUserData))
	}

	// Generate a JWT for the user
	jwt, err := utils.GenerateJWT(user)
	if err != nil {
		// Return error if generating the JWT token fails
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToGenerateTokenAccess))
	}

	// Send a greeting email to the user
	err = mail.SendGreetingEmail(user.Email, user.Username)
	if err != nil {
		// Return error if sending the greeting email fails
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.GreetingEmailError))
	}

	// Return success response with user data and token
	return output.GetSuccess(c, fiber.Map{
		"message": "Register account successful!",
		"data": fiber.Map{
			"user": user,
		},
		"token": jwt,
	})

}

// Login handles user login
func Login(c *fiber.Ctx) error {

	// Create a context with timeout for database operations
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Parse the incoming request body into a Login struct
	var login data.Login
	err := c.BodyParser(&login)
	if err != nil {
		// Return error if parsing the login data fails
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	// Validate the structure of the login object
	err = global.GetValidator().Struct(login)
	if err != nil {
		// Return error if the login data is invalid
		return output.GetError(c, fiber.StatusBadRequest, string(constant.ValidationError))
	}

	// Define filter to find user by email or username
	filter := bson.M{
		"$or": []bson.M{
			{"email": login.UserIdentifier},
			{"username": login.UserIdentifier},
		},
	}

	// Retrieve the user from the database
	user, err := helper.CheckUser(ctx, filter)
	if err != nil {
		// Return error if no account is found
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidAccountError))
	}

	// Check if the provided password matches the stored password
	err = helper.CheckPassword(user.Password, login.Password)
	if err != nil {
		// Return error if the password does not match
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidAccountError))
	}

	// Generate a JWT for the user
	jwt, err := utils.GenerateJWT(user)
	if err != nil {
		// Return error if generating the JWT token fails
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToGenerateTokenAccess))
	}

	// Return success response with user data and token
	return output.GetSuccess(c, fiber.Map{
		"message": "Login successful!",
		"data": fiber.Map{
			"user": user,
		},
		"token": jwt,
	})

}