package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atla/lotd/game"
	"github.com/atla/lotd/users"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan WebSocketMessage)  // broadcast channel
// Configure the upgrader
var upgrader = websocket.Upgrader{}

// WebSocketServer ... Define our message object
type WebSocketServer struct {
	port string
	game *game.Game
}

// NewWebSocketServer ... creates a new websocketserver instance
func NewWebSocketServer(port string) *WebSocketServer {
	return &WebSocketServer{
		port: port,
		game: game.GetInstance(),
	}
}

// WebSocketMessage ... Define our message object
type WebSocketMessage struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

// NewWebSocketMessage ... creates a new Websocket message
func NewWebSocketMessage(user string, message string) *WebSocketMessage {
	return &WebSocketMessage{
		Username: user,
		Message:  message,
	}
}

func (server *WebSocketServer) handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	ws.WriteJSON(NewWebSocketMessage("", server.game.MOTD))

	for {
		var msg WebSocketMessage
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		var user = users.GetInstance().FindUserByID(msg.Username)

		// send join message
		if !user.Active {
			user.Active = true
			server.game.OnUserJoined <- game.NewUserJoined(user)
		}

		var message = game.NewMessage(user, msg.Message)

		server.game.OnMessageReceived <- message
	}
}

func (server *WebSocketServer) handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast

		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// OnMessage .. broadcast receiver
func (server *WebSocketServer) OnMessage(message *game.Message) {

	fmt.Println("WebSocket Server received message")

	var userName string = ""
	if message.FromUser != nil {
		userName = message.FromUser.ID
	}

	// only respond to the target user
	if message.ToUser != nil {

	} else {
		// else broadcast this message
		broadcast <- WebSocketMessage{
			Username: userName,
			Message:  message.Data,
		}
	}

}

// OnSystemMessage .. broadcast receiver
func (server *WebSocketServer) OnSystemMessage(message *game.Message) {

	broadcast <- WebSocketMessage{
		Username: "#SYSTEM",
		Message:  message.Data,
	}
}

// Start ... start websocketserver
func (server *WebSocketServer) Start() {

	// Configure websocket route
	http.HandleFunc("/ws", server.handleConnections)

	// Start listening for incoming chat messages
	go server.handleMessages()

	game.GetInstance().Subscribe(server)

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :" + server.port)
	err := http.ListenAndServe(":"+server.port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
