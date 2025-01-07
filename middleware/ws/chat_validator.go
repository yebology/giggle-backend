package ws

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/constant"
	"github.com/yebology/giggle-backend/middleware/helper"
	"github.com/yebology/giggle-backend/output"
)

func ValidateChatSender(c *fiber.Ctx) error {

	claims, err := helper.ParseToken(c)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidTokenError))
	}

	var expectedSenderId = c.Query("senderId")
	senderId, ok := claims["id"].(string)
		if senderId != expectedSenderId || !ok {
			return output.GetError(c, fiber.StatusBadRequest, string(constant.PermissionDeniedError))
		}

	return c.Next()
	
}