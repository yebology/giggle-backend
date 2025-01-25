package ws

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/constant"
	"github.com/yebology/giggle-backend/middleware/helper"
	"github.com/yebology/giggle-backend/output"
)

// ValidateChatSender is a middleware that checks whether the sender of a chat message is authorized to send the message.
// It ensures that the sender's ID in the JWT token matches the "senderId" provided in the query parameters.
func ValidateChatSender(c *fiber.Ctx) error {

	// Parse the JWT token to extract the claims (which include the user information such as user ID).
	claims, err := helper.ParseToken(c)
	if err != nil {
		// If parsing the token fails, return a "Bad Request" error with an "Invalid Token" message.
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidTokenError))
	}

	// Get the "senderId" query parameter from the request to validate the sender.
	var expectedSenderId = c.Query("senderId")

	// Extract the user ID (sender) from the claims and check if it matches the expected sender ID from the query.
	senderId, ok := claims["id"].(string)
	if senderId != expectedSenderId || !ok {
		// If the sender IDs do not match or the user ID is not found in the claims, return a "Permission Denied" error.
		return output.GetError(c, fiber.StatusBadRequest, string(constant.PermissionDeniedError))
	}

	// If the sender is valid, proceed to the next middleware or handler.
	return c.Next()
	
}
