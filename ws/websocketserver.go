package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/atla/lotd/game"
	"github.com/atla/lotd/users"

	"github.com/gorilla/websocket"
)

// Server ... Define our message object
type Server struct {
	port      string
	game      *game.Game
	Clients   map[*websocket.Conn]bool
	Users     map[string]*websocket.Conn
	Broadcast chan Message
	upgrader  websocket.Upgrader

	MessageHandler *MessageHandler
}

// NewServer ... creates a new websocketserver instance
func NewServer(port string) *Server {
	server := &Server{
		port:      port,
		Clients:   make(map[*websocket.Conn]bool),
		Users:     make(map[string]*websocket.Conn),
		Broadcast: make(chan Message),
		upgrader:  websocket.Upgrader{},
		game:      game.GetInstance(),
	}

	server.MessageHandler = NewMessageHandler(server)

	return server
}

// Message ... Define our message object
type Message struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

// NewWebSocketMessage ... creates a new Websocket message
func NewWebSocketMessage(user string, message string) *Message {
	return &Message{
		Type:     "message",
		Username: user,
		Message:  message,
	}
}

// DisplayRoom ...
type DisplayRoom struct {
	Type string     `json:"type"`
	Room *game.Room `json:"room"`
}

// NewDisplayRoom ... creates new display room message
func NewDisplayRoom(room *game.Room) *DisplayRoom {

	if room == nil {
		log.Panic("no room to create displayRoom message")
	}

	return &DisplayRoom{
		Type: "displayRoom",
		Room: room,
	}
}

func (server *Server) handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	server.Clients[ws] = true

	ws.WriteJSON(NewWebSocketMessage("", server.game.MOTD))

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(server.Clients, ws)
			break
		}

		user, _ := users.GetInstance().FindUserByID(msg.Username)

		// send join message
		if !user.Active {
			server.Users[user.ID] = ws
			user.Active = true
			server.game.OnUserJoined <- game.NewUserJoined(user)
		}

		if msg.Message != "" {
			var message = game.NewMessage(user, msg.Message)
			server.game.OnMessageReceived <- message
		}
	}
}

func (server *Server) sendMessage(id string, msg interface{}) {

	switch x := msg.(type) {
	case Message:
		fmt.Println(json.MarshalIndent(x, "", "    "))

	case DisplayRoom:
		fmt.Println(json.MarshalIndent(x, "", "    "))

	}

	if client, ok := server.Users[id]; ok {
		err := client.WriteJSON(msg)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			delete(server.Clients, client)
			delete(server.Users, id)
		}
	}
}

func (server *Server) handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-server.Broadcast

		// Send it out to every client that is currently connected
		for client := range server.Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(server.Clients, client)
			}
		}
	}
}

// OnSystemMessage .. broadcast receiver
func (server *Server) OnSystemMessage(message *game.Message) {

	server.Broadcast <- Message{
		Username: "#SYSTEM",
		Message:  message.Data,
	}
}

// Start ... start websocketserver
func (server *Server) Start() {

	// Configure websocket route
	http.HandleFunc("/ws", server.handleConnections)

	// Start listening for incoming chat messages
	go server.handleMessages()

	game.GetInstance().Subscribe(server.MessageHandler)

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :" + server.port)
	err := http.ListenAndServe(":"+server.port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
