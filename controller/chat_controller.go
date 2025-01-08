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

type Hub struct {

	Clients						map[primitive.ObjectID]*websocket.Conn
	ClientRegisterChannel		chan *websocket.Conn
	ClientRemovalChannel		chan *websocket.Conn
	BroadcastChat				chan ws.Chat
	Group						chan http.Group
	BroadcastGroupChat			chan ws.GroupChat

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

		case groupMsg := <- h.BroadcastGroupChat:

			for _, receiverId := range groupMsg.ReceiverIds {
				receiverConn, ok := h.Clients[receiverId]
				if ok {
					receiverConn.WriteJSON(groupMsg.Chat)
				}
			}

		}

	}

}

func Chat(h *Hub) func (*websocket.Conn) {

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

				senderObjectId, receiverObjectId, err := helper.ConvertToObjectIdBoth(senderId, receiverId)
				if err != nil {
					log.Println("Error converting senderId to objectId:", err)
					return
				}

				isGroupChatStr := conn.Query("isGroupChat")
				chatType, err := helper.CheckChatType(isGroupChatStr)
				if err != nil {
					log.Println("Error converting chat type data:", err)
					return
				}

				chat := ws.Chat{
					SenderId: senderObjectId,
					ReceiverId: receiverObjectId,
					Message: string(message),
					ChatTimestamp: uint64(time.Now().Unix()),
					ChatType: chat.Chat(chatType),
				}

				collection := database.GetDatabase().Collection("chat")
				_, err = collection.InsertOne(context.Background(), chat)
				if err != nil {
					log.Println("Error while sending a message:", err)
					return
				}				

				if chatType == "Group" {

					filter := bson.M{"_groupId": receiverId}

					receiverIds, err := helper.GetGroupUsersId(filter)
					if err != nil {
						log.Println("Error while get group users:", err)
						return
					}

					groupChat := ws.GroupChat{
						Chat: chat,
						ReceiverIds: receiverIds,
					}

					h.BroadcastGroupChat <- groupChat

				} else {

					h.BroadcastChat <- chat

				}

			}


		}

	}

}