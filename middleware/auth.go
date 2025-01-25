package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/constant"
	"github.com/yebology/giggle-backend/middleware/helper"
	"github.com/yebology/giggle-backend/output"
)

// AuthMiddleware is a middleware function that authenticates the user by verifying the JWT token.
// It ensures the user is authorized to access the requested route based on their role.
func AuthMiddleware(c *fiber.Ctx) error {

	// Parse the JWT token from the request to extract the claims (payload).
	// If there's an error or the token is invalid, return an Unauthorized error response.
	claims, err := helper.ParseToken(c)
	if err != nil {
		return output.GetError(c, fiber.StatusUnauthorized, string(constant.InvalidTokenError))
	}

	// Define the expected role for the user to access the route (in this case, "user").
	var expectedRole = "user"

	// Extract the user's role from the claims to verify if they have the appropriate access.
	// If the role doesn't match or if there is an issue extracting the role, return a Forbidden error.
	role, ok := claims["role"].(string)
	if role != expectedRole || !ok {
		return output.GetError(c, fiber.StatusForbidden, string(constant.PermissionDeniedError))
	}

	// If the token is valid and the role matches, allow the request to proceed to the next middleware or handler.
	return c.Next()
	
}
