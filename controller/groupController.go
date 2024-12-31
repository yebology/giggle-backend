package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/output"
)

func GetGroup(c *fiber.Ctx) error {

	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully fetch user's group!",
		"data": "",
	})

}