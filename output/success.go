package output

import "github.com/gofiber/fiber/v2"

// GetSuccess is a helper function that sends a success response in JSON format.
// It sets the HTTP status code to 200 (OK) and returns the provided data in the response body.
func GetSuccess(c *fiber.Ctx, data fiber.Map) error {

	// Return a JSON response with status 200 (OK) and the provided success data.
	return c.Status(fiber.StatusOK).JSON(data)
	
}
