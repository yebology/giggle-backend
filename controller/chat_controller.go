package controller

import (
	"context"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/yebology/giggle-backend/controller/helper"
	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/model/ws"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hub struct {

	Clients						map[primitive.ObjectID]*websocket.Conn
	ClientRegisterChannel		chan *websocket.Conn
	ClientRemovalChannel		chan *websocket.Conn
	BroadcastChat				chan ws.Chat
	Groups						map[primitive.ObjectID][]primitive.ObjectID

}

func (h *Hub) Run() {
	
	for {

		select {

		case conn := <- h.ClientRegisterChannel:

			senderId := conn.Query("senderId")

			res, err := helper.ConvertToObjectId(senderId)
			if err != nil {
				log.Println("Error converting senderId to objectId:", err)
				return
			}
			
			h.Clients[res] = conn

		case conn := <- h.ClientRemovalChannel:

			senderId := conn.Query("senderId")

			res, err := helper.ConvertToObjectId(senderId)
			if err != nil {
				log.Println("Error converting senderId to objectId:", err)
				return
			}

			delete(h.Clients, res)

		case msg := <- h.BroadcastChat:

			receiverConn, ok := h.Clients[msg.ReceiverId]
			if ok {
				receiverConn.WriteJSON(msg)
			}

		}

	}

}

func PersonalChat(h *Hub) func (*websocket.Conn) {

	return func(conn *websocket.Conn) {

		defer func() {

			h.ClientRemovalChannel <- conn
			conn.Close()

		}()

		h.ClientRegisterChannel <- conn

		for {

			messageType, message, err := conn.ReadMessage() 
			if err != nil {
				log.Println("Error whiler register a new client connection:", err)
				return
			}
			if messageType == websocket.TextMessage {

				senderId := conn.Query("senderId")
				receiverId := conn.Query("receiverId")

				senderObjectId, err := helper.ConvertToObjectId(senderId)
				if err != nil {
					log.Println("Error converting senderId to objectId:", err)
					return
				}

				receiverObjectId, err := helper.ConvertToObjectId(receiverId)
				if err != nil {
					log.Println("Error converting receiverId to objectId:", err)
					return
				}

				chat := ws.Chat{
					SenderId: senderObjectId,
					ReceiverId: receiverObjectId,
					Message: string(message),
				}

				collection := database.GetDatabase().Collection("chat")
				_, err = collection.InsertOne(context.Background(), chat)
				if err != nil {
					log.Println("Error while sending a message:", err)
				} else {
					h.BroadcastChat <- chat
				}

			}


		}

	}

}