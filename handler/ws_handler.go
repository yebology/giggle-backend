package handler

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/controller"
	"github.com/yebology/giggle-backend/middleware"
	WsMiddleware "github.com/yebology/giggle-backend/middleware/ws"
	"github.com/yebology/giggle-backend/model/http"
	"github.com/yebology/giggle-backend/model/ws"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SetUp configures the routes and WebSocket handlers for the application.
// It also initializes the WebSocket hub and middleware.
func SetUp(app *fiber.App) {

	// Initialize the WebSocket hub with channels to manage client connections and messages
	hub := &controller.Hub{
		// Map to keep track of connected clients using their ObjectID
		Clients: make(map[primitive.ObjectID]*websocket.Conn),
		
		// Channels to handle client registration, removal, and messages
		ClientRegisterChannel: make(chan *websocket.Conn),
		ClientRemovalChannel:  make(chan *websocket.Conn),
		
		// Channels for broadcasting chat messages and group messages
		BroadcastChat:        make(chan ws.Chat),
		Group:                make(chan http.Group),
		BroadcastGroupChat:   make(chan ws.GroupChat),
	}

	// Start the hub's message handling process in a separate goroutine
	go hub.Run()

	// Set up WebSocket middleware for the "/ws" endpoint
	// - AuthMiddleware: Authenticates the user before proceeding
	// - ValidateChatSender: Validates the sender of the chat message
	// - ValidateWebSocketUpgrade: Ensures the WebSocket upgrade request is valid
	app.Use("/ws", middleware.AuthMiddleware, WsMiddleware.ValidateChatSender, WsMiddleware.ValidateWebSocketUpgrade)

	// Define the WebSocket route for chat, linking it to the Chat handler
	// The hub is passed as a parameter to the handler
	app.Get("/ws/chat", websocket.New(controller.Chat(hub)))
	
}
