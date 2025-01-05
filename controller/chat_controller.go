package controller

import (
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/yebology/giggle-backend/model/ws"
)

type Hub struct {

	Clients						map[*websocket.Conn]bool
	ClientRegisterChannel		chan *websocket.Conn
	ClientRemovalChannel		chan *websocket.Conn
	BroadcastChat				chan ws.PersonalChat

}

func (h *Hub) Run() {
	
	for {

		select {

		case conn := <- h.ClientRegisterChannel:
			h.Clients[conn] = true

		case conn := <- h.ClientRemovalChannel:
			delete(h.Clients, conn)

		case msg := <- h.BroadcastChat:
			for conn := range h.Clients {

				conn.WriteJSON(msg)

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

		senderId := conn.Query("senderId")
		receiverId := conn.Query("receiverId")
		h.ClientRegisterChannel <- conn

		for {

			messageType, message, err := conn.ReadMessage() 
			if err != nil {
				fmt.Println(err)
				return
			}
			if messageType == websocket.TextMessage {
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