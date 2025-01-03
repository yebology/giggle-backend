package ws

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/router/ws/helper"
)

func SetUp(app *fiber.App) {

	app.Use("/ws", helper.ValidateWebSocketUpgrade)

	app.Get("/ws/:id", websocket.New(helper.HandleWebSocket))

}