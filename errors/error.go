package errors

import "github.com/gofiber/fiber/v2"

// type Error int
// const (
// 	BodyParse Error = iota
// 	HashError

// )

func GetError(c *fiber.Ctx, err string) error {

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":err})
	
}