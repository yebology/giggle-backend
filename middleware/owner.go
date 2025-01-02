package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/constant"
	"github.com/yebology/giggle-backend/middleware/helper"
	"github.com/yebology/giggle-backend/output"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PostOwnerMiddleware(c *fiber.Ctx) error {

	claims, err := helper.ParseToken(c)
	if err != nil {
		return output.GetError(c, fiber.StatusUnauthorized, string(constant.InvalidTokenError))
	}

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)


	return c.Next()

}