package output

import "github.com/gofiber/fiber/v2"

func GetSuccess(c *fiber.Ctx, data fiber.Map) error {

	return c.Status(fiber.StatusOK).JSON(data)

}