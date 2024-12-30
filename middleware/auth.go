package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/output"
	"github.com/yebology/giggle-backend/model/constant"
)

func UserMiddleware(c *fiber.Ctx) error {

	role := GetRoleFromContext(c)
	if role != string(constant.User) {
		return output.GetError(c, fiber.StatusForbidden, "Permission denied! Must register or login first!")
	}

	return c.Next()

}

func GetRoleFromContext(c *fiber.Ctx) string {
	
	return ""

}