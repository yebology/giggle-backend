package helper

import (
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func ValidateWebSocketUpgrade(c *fiber.Ctx) error {

	if websocket.IsWebSocketUpgrade(c) {

		c.Locals("allowed", true)

		return c.Next()

	}

	return fiber.ErrUpgradeRequired

}

func HandleWebSocket(c *websocket.Conn) {

	allowed := c.Locals("allowed")
	if allowed != true {
		fmt.Println(allowed)
		return
	}

	fmt.Println("ID: ", c.Params("id"))
	fmt.Println("Version: ", c.Query("v"))
	fmt.Println("Session: ", c.Cookies("session"))

	for {

		mt, msg, err := c.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return 
		}

		fmt.Println(msg)

		err = c.WriteMessage(mt, msg)
		if err != nil {
			fmt.Println(err)
			return 
		}

	}

}