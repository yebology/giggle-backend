package output

import (
	"github.com/gofiber/fiber/v2"
)

// GetError is a helper function that sends an error response in JSON format.
// It sets the HTTP status code and returns an error message in the response body.
func GetError(c *fiber.Ctx, status int, err string) error {

	// Return a JSON response with the given status code and error message.
	return c.Status(status).JSON(fiber.Map{
		"message": err, // The error message to be returned.
		"data": "",     // Empty data field, as the focus is on returning the error message.
	})
	
}
