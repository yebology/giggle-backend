package ws

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

// ValidateWebSocketUpgrade is a middleware that checks if the incoming request is a valid WebSocket upgrade request.
// It ensures that the client is attempting to establish a WebSocket connection, and if so, it allows the connection to proceed.
func ValidateWebSocketUpgrade(c *fiber.Ctx) error {

	// Check if the current request is a valid WebSocket upgrade request.
	if websocket.IsWebSocketUpgrade(c) {

		// If it is a valid WebSocket upgrade, set a local variable to indicate the connection is allowed.
		// This can be used later in the request handling to verify the WebSocket upgrade is authorized.
		c.Locals("allowed", true)

		// Proceed to the next middleware or handler in the request pipeline.
		return c.Next()
	}

	// If the request is not a valid WebSocket upgrade, return an error indicating that upgrading to WebSocket is required.
	return fiber.ErrUpgradeRequired
	
}
