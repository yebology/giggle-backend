package middleware

import "github.com/gofiber/fiber/v2"

func PostOwnerMiddleware(c *fiber.Ctx) error {

	return c.Next()

}