package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	game "github.com/atla/lotd/game"
	"github.com/atla/lotd/login"
	"github.com/atla/lotd/users"
)

// Client ... asd
type Client struct {
	incoming      chan string
	outgoing      chan string
	reader        *bufio.Reader
	writer        *bufio.Writer
	connection    net.Conn
	authenticated bool
	user          *users.User
}

// NewClient ... creates a new client
func NewClient(connection net.Conn) *Client {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	client := &Client{
		incoming:      make(chan string),
		outgoing:      make(chan string),
		reader:        reader,
		writer:        writer,
		connection:    connection,
		authenticated: false,
	}

	client.Listen()

	return client
}

func (client *Client) Read() {
	for {
		line, _ := client.reader.ReadString('\n')

		if len(line) > 1 {
			line = line[0 : len(line)-1]
		}

		client.incoming <- line
	}
}

func (client *Client) Write() {
	for data := range client.outgoing {
		client.writer.WriteString(data)
		client.writer.Flush()
	}
}

// Listen ... asda
func (client *Client) Listen() {
	go client.Read()
	go client.Write()
}

// Broadcast ...
func (server *Server) Broadcast(data string) {
	for _, client := range server.clients {
		client.outgoing <- data
	}
}

func (client *Client) loginWithPassword() bool {

	var userManager = users.GetInstance()

	//	var echoOff = []byte{0xFF, 0xFC, 0x01, 0x0}
	//	var echoOn = []byte{0xFF, 0xFB, 0x01, 0x0}

	log.Println("User logging in")

	client.outgoing <- "User: "

	var user = <-client.incoming
	user = strings.TrimSuffix(strings.TrimSuffix(user, "\n"), "\r")

	// turn of user echo on client side
	//client.connection.Write(echoOff)

	//client.outgoing <- string(echoOff)

	client.outgoing <- "Password: "
	var password = <-client.incoming
	password = strings.TrimSuffix(strings.TrimSuffix(password, "\n"), "\r")

	// turn on user echo on client side
	//	client.outgoing <- string(echoOn)

	var loginManager = login.NewLoginManager()
	var loginSuccessful = loginManager.Login(user, password)

	if loginSuccessful {
		client.user, _ = userManager.FindUserByID(user)

		fmt.Println("FOUND USER " + client.user.ID)

		return true
	}
	return false
}

func (client *Client) registerNewUser() bool {

	log.Println("Creating new user")

	var userExists = true
	var user = ""

	for userExists {
		client.outgoing <- "User: "
		user = <-client.incoming
		user = strings.TrimSuffix(strings.TrimSuffix(user, "\n"), "\r")

		if _, err := users.GetInstance().FindUserByID(user); err != nil {
			userExists = false
		} else {
			log.Println("User already exists.")
		}
	}

	var passwordMatches = false
	var password = ""

	for !passwordMatches {
		client.outgoing <- "Password: "
		password = <-client.incoming
		password = strings.TrimSuffix(strings.TrimSuffix(password, "\n"), "\r")

		client.outgoing <- "Password (repeat): "
		var password2 = <-client.incoming
		password2 = strings.TrimSuffix(strings.TrimSuffix(password2, "\n"), "\r")

		passwordMatches = password == password2

		if !passwordMatches {
			client.outgoing <- "Passwords do not match"
		}
	}

	var userManager = users.GetInstance()
	var newUser = users.NewUser(user, password, "tcp_signup")
	userManager.AddUser(newUser)

	fmt.Println("Created new user " + user)

	client.user = newUser

	return true
}

// Join ... handles a new client joining
func (server *Server) Join(connection net.Conn) {
	client := NewClient(connection)

	client.outgoing <- server.game.MOTD
	client.outgoing <- "Welcome to the Lair of the Dragon\n(1) Existing Account (use guest:guest to look around)\n(2) New Account\nChoose: "

	// handle first account choice

	go func() {

		var loginPassed = false

		for !loginPassed {
			var choice = <-client.incoming

			if strings.HasPrefix(choice, "1") {
				loginPassed = client.loginWithPassword()
			} else if strings.HasPrefix(choice, "2") {
				loginPassed = client.registerNewUser()
			} else {
				client.outgoing <- "Welcome to the Lair of the Dragon\n(1) Existing Account\n(2) New Account\nChoose: "
			}

		}

		// add the client to the list of current active clients
		server.clients = append(server.clients, client)

		client.outgoing <- "Connected.\n"

		server.game.OnUserJoined <- game.NewUserJoined(client.user)

		for {

			var clientMessage = <-client.incoming

			if clientMessage != "" {

				log.Println("TCP: " + clientMessage)

				// TODO: exit should be a command?
				if strings.HasPrefix(clientMessage, "exit") {
					server.onClientQuit(client)
					return
				}

				log.Println("Forwarding message to game instance " + clientMessage)

				server.game.OnMessageReceived <- game.NewMessage(client.user, clientMessage)

			}
		}
	}()
}

// Server ... tbd
type Server struct {
	port     string
	clients  []*Client
	joins    chan net.Conn
	incoming chan string
	outgoing chan string

	game           *game.Game
	MessageHandler *MessageHandler
}

func (server *Server) onClientQuit(client *Client) {
	for index, c := range server.clients {
		if c == client {
			server.clients = append(server.clients[:index], server.clients[index+1:]...)
			client.connection.Close()

			server.game.OnUserQuit <- game.NewUserQuit(client.user)

			return
		}
	}
}

// NewServer ... creates and returns a new TCPGameServer instance
func NewServer(port string) *Server {

	server := &Server{
		clients:  make([]*Client, 0),
		joins:    make(chan net.Conn),
		incoming: make(chan string),
		outgoing: make(chan string),
		game:     game.GetInstance(),
		port:     port,
	}

	server.MessageHandler = NewMessageHandler(server)
	server.game.Subscribe(server.MessageHandler)
	server.listen()

	return server
}

func (server *Server) getClientByID(id string) (*Client, bool) {
	for _, client := range server.clients {
		if client.user.ID == id {
			return client, true
		}
	}
	// not found
	return nil, false
}

// Start .. starts the created server
func (server *Server) Start() {

	listener, _ := net.Listen("tcp", ":"+server.port)
	log.Println("Started TCP server on port " + server.port)
	for {
		conn, _ := listener.Accept()
		server.joins <- conn
	}
}

// Run ... processes every entity
func (server *Server) listen() {

	go func() {
		for {
			select {
			//case data := <-server.incoming:

			//				server.Broadcast(data)
			case conn := <-server.joins:
				server.Join(conn)
			}
		}
	}()
}
