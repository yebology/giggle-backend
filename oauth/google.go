package oauth

import (
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/constant"
	"github.com/yebology/giggle-backend/controller/helper"
	"github.com/yebology/giggle-backend/global"
	"github.com/yebology/giggle-backend/model/data"
	"github.com/yebology/giggle-backend/output"
	"github.com/yebology/giggle-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// Generates the Google OAuth authentication URL and redirects the user to it for login.
func GoogleAuth(c *fiber.Ctx) error {

	// Get authentication URL.
	url := global.GetGoogleOauth().AuthCodeURL("state")

	// Redirect the user to Google login page.
	return c.Redirect(url)

}

// Handles the Google OAuth redirect, exchanges the authorization code for an access token,
// retrieves the user's information, and returns a success response.
func GoogleRedirect(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get code from redirect url.
	code := c.Query("code")
	if code == "" {
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToGetCodeFromRedirectUrl))
	}

	// Exchange the code with an access token.
	token, err := global.GetGoogleOauth().Exchange(context.Background(), code)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToExchangeCodeWithToken))
	}

	// Get user information by using access token.
	client := global.GetGoogleOauth().Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToLoadUserData))
	}

	// Read the response body to extract user data.
	// Ensure the response body is closed after reading to release resources and maintain security.
	defer response.Body.Close()
	var googleUser data.GoogleUser
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToLoadUserData))
	}

	// Parse/decode the JSON response and map it into 'googleUser' structure.
	err = json.Unmarshal(bytes, &googleUser)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToDecodeData))
	}

	// Filter the 'User' by email then return the result.
	filter := bson.M{"email": googleUser.Email}
	user, err := helper.CheckUser(ctx, filter)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.UnregisteredAccountError))
	}

	// Generate JWT token to handle 'User' authorization.
	jwt, err := utils.GenerateJWT(user)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToGenerateTokenAccess))
	}

	// Send a success response.
	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully login using Google account!",
		"data": fiber.Map{
			"user": user,
		},
		"token": jwt,
	})

}