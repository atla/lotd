package tcp

import (
	"fmt"
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

	log.Println("tcpmessagehandler OnMessage")

	s := messageHandler.server

	switch msg := message.(type) {
	case *game.Message:

		fmt.Println("Received message")

		var msgstring = ""
		// message with user context
		if msg.FromUser != nil {
			msgstring = msg.FromUser.ID + ": " + msg.Data + "\n"
		} else {
			msgstring = msg.Data + "\n"
		}

		//TODO: dont send message to own client

		if msg.ToUser != nil {
			for _, client := range messageHandler.server.clients {
				if client.user.ID == msg.ToUser.ID {
					client.outgoing <- msgstring
				}
			}
		} else {
			messageHandler.server.Broadcast(msgstring)
		}

	case *game.AvatarJoinedRoom:

		if client, ok := s.getClientByID(msg.Avatar.CurrentUser.ID); ok {
			client.outgoing <- "=== " + msg.Room.Title + " ===\n" + msg.Room.Description
		}

		for _, avatar := range msg.Room.Avatars {

			if avatar != msg.Avatar {
				if client, ok := messageHandler.server.getClientByID(avatar.ID); ok {
					client.outgoing <- avatar.ID + " appeared."
				}
			}
		}

	case *game.AvatarLeftRoom:
		fmt.Println("Unhandled")
	default:
		log.Println("Unknown message type in tcpmessagehandler ")

	}

}
