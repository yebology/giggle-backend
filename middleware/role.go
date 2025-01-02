package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/constant"
	"github.com/yebology/giggle-backend/middleware/helper"
	"github.com/yebology/giggle-backend/output"
)

func UserMiddleware(c *fiber.Ctx) error {

	claims, err := helper.ParseToken(c)
	if err != nil {
		return output.GetError(c, fiber.StatusUnauthorized, string(constant.InvalidTokenError))
	}

	var expectedRole = "user"
	role, ok := claims["role"].(string)
	if role != expectedRole || !ok {
		return output.GetError(c, fiber.StatusForbidden, string(constant.PermissionDeniedError))
	}

	return c.Next()

}
