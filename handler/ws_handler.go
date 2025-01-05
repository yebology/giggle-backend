package handler

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/controller"
	"github.com/yebology/giggle-backend/middleware"
	"github.com/yebology/giggle-backend/model/ws"
)

func SetUp(app *fiber.App) {

	hub := &controller.Hub{

		Clients: make(map[*websocket.Conn]bool),
		ClientRegisterChannel: make(chan *websocket.Conn),
		ClientRemovalChannel: make(chan *websocket.Conn),
		BroadcastChat: make(chan ws.PersonalChat),

	}

	go hub.Run()

	app.Use("/ws", middleware.ValidateWebSocketUpgrade)

	app.Use("/ws/personalChat", websocket.New(controller.PersonalChat(hub)))

	app.Use("/ws/groupChat", websocket.New(controller.GroupChat))

}