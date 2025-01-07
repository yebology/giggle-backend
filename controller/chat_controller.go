package controller

import (
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/yebology/giggle-backend/model/ws"
)

type Hub struct {

	Clients						map[string]*websocket.Conn
	ClientRegisterChannel		chan *websocket.Conn
	ClientRemovalChannel		chan *websocket.Conn
	BroadcastChat				chan ws.PersonalChat

}

func (h *Hub) Run() {
	
	for {

		select {

		case conn := <- h.ClientRegisterChannel:
			senderId := conn.Query("senderId")
			h.Clients[senderId] = conn

		case conn := <- h.ClientRemovalChannel:
			senderId := conn.Query("senderId")
			delete(h.Clients, senderId)

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
				fmt.Println(err)
				return
			}
			if messageType == websocket.TextMessage {

				senderId := conn.Query("senderId")
				receiverId := conn.Query("receiverId")
				chat := ws.PersonalChat{
					SenderId: senderId,
					ReceiverId: receiverId,
					Message: string(message),
				}
				h.BroadcastChat <- chat

			}


		}

	}

}

func GroupChat(c *websocket.Conn) {



}