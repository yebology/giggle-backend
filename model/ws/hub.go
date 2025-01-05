package ws

import (
	"github.com/gofiber/contrib/websocket"
)

type Hub struct {

	Clients						map[*websocket.Conn]bool
	ClientRegisterChannel		chan *websocket.Conn
	ClientRemovalChannel		chan *websocket.Conn
	BroadcastChat				chan PersonalChat

}