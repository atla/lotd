package ws

import (
	"log"

	game "github.com/atla/lotd/game"
)

// MessageHandler ... tcp ..
type MessageHandler struct {
	server *Server
}

// NewMessageHandler ... creates a new message handler
func NewMessageHandler(server *Server) *MessageHandler {
	return &MessageHandler{
		server: server,
	}
}

// OnMessage .. broadcast receiver
func (messageHandler *MessageHandler) OnMessage(message interface{}) {

	log.Println("websocketmessagehandler OnMessage")

	s := messageHandler.server

	switch msg := message.(type) {
	case *game.Message:

		var userName string

		if msg.FromUser != nil {
			userName = msg.FromUser.ID
		}

		// only respond to the target user
		if msg.ToUser != nil {

		} else {
			// else broadcast this message
			s.Broadcast <- Message{
				Username: userName,
				Message:  msg.Data,
			}
		}

	case *game.AvatarJoinedRoom:

		for _, avatar := range msg.Room.Avatars {
			if avatar != msg.Avatar {

				s.Broadcast <- Message{
					Message: msg.Avatar.ID + " appeared.",
				}

			}
		}

		s.sendMessage(msg.Avatar.ID, NewDisplayRoom(msg.Room))

	case *game.AvatarLeftRoom:

	}

}
