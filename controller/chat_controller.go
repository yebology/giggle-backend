package controller

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/yebology/giggle-backend/constant/chat"
	"github.com/yebology/giggle-backend/controller/helper"
	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/model/http"
	"github.com/yebology/giggle-backend/model/ws"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Hub represents the WebSocket hub that manages client connections, message broadcasting, and group chat.
type Hub struct {
	// Clients holds all active WebSocket connections indexed by their sender ID
	Clients map[primitive.ObjectID]*websocket.Conn

	// Channels for registering and removing WebSocket clients, broadcasting messages
	ClientRegisterChannel chan *websocket.Conn
	ClientRemovalChannel  chan *websocket.Conn

	// Channels for chat messages and group messages
	BroadcastChat        chan ws.Chat
	Group                chan http.Group
	BroadcastGroupChat   chan ws.GroupChat
}

// Run starts the hub to handle client connections and broadcast messages
func (h *Hub) Run() {

	for {
		select {
			
		// Register new client connection
		case conn := <-h.ClientRegisterChannel:

			// Convert sender ID from query string to ObjectID
			senderId := conn.Query("senderId")
			res, err := helper.ConvertToObjectId(senderId)
			if err != nil {
				log.Println("Error converting senderId to objectId:", err)
				return
			}
			h.Clients[res] = conn

		// Remove client connection
		case conn := <-h.ClientRemovalChannel:

			// Convert sender ID from query string to ObjectID
			senderId := conn.Query("senderId")
			res, err := helper.ConvertToObjectId(senderId)
			if err != nil {
				log.Println("Error converting senderId to objectId:", err)
				return
			}
			delete(h.Clients, res)

		// Handle a regular chat message to be broadcasted
		case msg := <-h.BroadcastChat:

			// Retrieve the WebSocket connection for the receiver
			receiverConn, ok := h.Clients[msg.ReceiverId]

			// Generate a key for decrypting the message
			key := helper.Generate32BytesKey(msg.SenderId, msg.ReceiverId)

			// Decrypt the message with AES256 encryption
			decMessage, err := helper.DecryptMessageWithAES256(msg.Message, key)
			if err != nil {
				log.Println("Error while decrypting message:", err)
				return
			}

			// If the receiver is connected, send the decrypted message
			if ok {
				receiverConn.WriteJSON(decMessage)
			}

		// Handle group chat message to be broadcasted
		case groupMsg := <-h.BroadcastGroupChat:

			// Generate a key for decrypting the group chat message
			key := helper.Generate32BytesKey(groupMsg.Chat.SenderId, groupMsg.Chat.ReceiverId)

			// Decrypt the message with AES256 encryption
			decMessage, err := helper.DecryptMessageWithAES256(groupMsg.Chat.Message, key)
			if err != nil {
				log.Println("Error while decrypting message:", err)
				return
			}

			// Send the decrypted message to all group members
			for _, receiverId := range groupMsg.ReceiverIds {

				// Skip sender to avoid sending the message back to the sender
				if receiverId != groupMsg.Chat.SenderId {
					receiverConn, ok := h.Clients[receiverId]
					if ok {
						receiverConn.WriteJSON(decMessage)
					}
				}

			}
		}
	}
}

// Chat handles a new WebSocket chat connection
func Chat(h *Hub) func(*websocket.Conn) {
	return func(conn *websocket.Conn) {

		// Ensure to remove the connection and close it when the function exits
		defer func() {
			h.ClientRemovalChannel <- conn
			conn.Close()
		}()

		// Register the new connection
		h.ClientRegisterChannel <- conn

		// Listen for incoming messages from the WebSocket connection
		for {

			// Read a message from the WebSocket
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error while register a new client connection:", err)
				return
			}

			// Only process text messages
			if messageType == websocket.TextMessage {

				// Extract sender and receiver IDs from query parameters
				senderId := conn.Query("senderId")
				receiverId := conn.Query("receiverId")

				// Convert sender and receiver IDs to ObjectIDs
				senderObjectId, receiverObjectId, err := helper.ConvertToObjectIdBoth(senderId, receiverId)
				if err != nil {
					log.Println("Error converting senderId to objectId:", err)
					return
				}

				// Check if this is a group chat
				isGroupChatStr := conn.Query("isGroupChat")
				chatType, err := helper.CheckChatType(isGroupChatStr)
				if err != nil {
					log.Println("Error converting chat type data:", err)
					return
				}

				// Generate a key for encrypting the message
				key := helper.Generate32BytesKey(senderObjectId, receiverObjectId)

				// Encrypt the message using AES256
				encMessage, err := helper.EncryptMessageWithAES256(string(message), key)
				if err != nil {
					log.Println("Error encrypting message:", err)
					return
				}

				// Create a new chat message
				chat := ws.Chat{
					SenderId:      senderObjectId,
					ReceiverId:    receiverObjectId,
					Message:       encMessage,
					ChatTimestamp: uint64(time.Now().Unix()),
					ChatType:      chat.Chat(chatType),
				}

				// Save the chat message to the database
				collection := database.GetDatabase().Collection("chat")
				_, err = collection.InsertOne(context.Background(), chat)
				if err != nil {
					log.Println("Error while sending a message:", err)
					return
				}

				// Handle group chat messages
				if chatType == "Group" {
					filter := bson.M{"_id": receiverObjectId}

					// Retrieve the group member IDs
					receiverIds, err := helper.GetGroupUsersId(filter)
					if err != nil {
						log.Println("Error while get group users:", err)
						return
					}

					// Create a GroupChat struct and broadcast the message
					groupChat := ws.GroupChat{
						Chat:        chat,
						ReceiverIds: receiverIds,
					}

					h.BroadcastGroupChat <- groupChat

				} else {
					// Broadcast the individual chat message
					h.BroadcastChat <- chat
				}
			}
		}
	}
}
