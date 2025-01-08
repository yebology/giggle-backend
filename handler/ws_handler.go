package handler

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/controller"
	"github.com/yebology/giggle-backend/middleware"
	WsMiddleware "github.com/yebology/giggle-backend/middleware/ws"
	"github.com/yebology/giggle-backend/model/ws"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetUp(app *fiber.App) {

	hub := &controller.Hub{

		Clients: 				make(map[primitive.ObjectID]*websocket.Conn),
		ClientRegisterChannel: 	make(chan *websocket.Conn),
		ClientRemovalChannel: 	make(chan *websocket.Conn),
		BroadcastChat: 			make(chan ws.Chat),
		Groups:					make(map[primitive.ObjectID][]primitive.ObjectID),

	}

	go hub.Run()

	app.Use("/ws", middleware.AuthMiddleware, WsMiddleware.ValidateChatSender, WsMiddleware.ValidateWebSocketUpgrade)

	app.Get("/ws/personalChat", websocket.New(controller.PersonalChat(hub)))

	// app.Get("/ws/groupChat/:id", websocket.New(controller.GroupChat))

}