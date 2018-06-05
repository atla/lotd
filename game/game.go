package game

import (
	"log"
	"lotd/users"
	"strings"
	"sync"
)

// CommandProcessor ... global user struct to control logins
type CommandProcessor struct {
	commands map[string]Command
}

// RegisterCommand ... register
func (commandProcessor *CommandProcessor) RegisterCommand(key string, command Command) {
	commandProcessor.commands[key] = command
}

// Process ...asd
func (commandProcessor *CommandProcessor) Process(game *Game, message *Message) bool {

	parts := strings.Fields(message.Data)

	if len(parts) > 0 {
		var key = parts[0]
		if val, ok := commandProcessor.commands[key]; ok {

			log.Println("Found command " + key + " executing...")
			return val.Execute(game, message)
		}
	}

	return false

}

// ScreamCommand ... foo
type ScreamCommand struct {
}

// Execute ... executes scream command
func (screamCommand *ScreamCommand) Execute(game *Game, message *Message) bool {

	parts := strings.Fields(message.Data)
	newMsg := strings.Join(parts[1:len(parts)], " ")

	var newMessage = "-- " + message.User.ID + " screams " + strings.ToUpper(newMsg) + "!!!!!"
	game.OnMessageReceived <- NewMessage(game.SystemUser, newMessage)
	return true

}

func (commandProcessor *CommandProcessor) registerCommands() {

	commandProcessor.RegisterCommand("scream", &ScreamCommand{})

}

// NewCommandProcessor .. creates a new command processor
func NewCommandProcessor() *CommandProcessor {
	var commandProcessor = &CommandProcessor{
		commands: make(map[string]Command),
	}
	// only once?
	commandProcessor.registerCommands()
	return commandProcessor
}

// Command ... commands
type Command interface {
	Execute(game *Game, message *Message) bool
}

//UserJoined ... player joined event
type UserJoined struct {
	User *users.User
}

//UserQuit ... player joined event
type UserQuit struct {
	User *users.User
}

// NewUserQuit ... creates a new User Joined event
func NewUserQuit(user *users.User) *UserQuit {
	return &UserQuit{
		User: user,
	}
}

// NewUserJoined ... creates a new User Joined event
func NewUserJoined(user *users.User) *UserJoined {
	return &UserJoined{
		User: user,
	}
}

// Message ... main message container to pass data from users to server and back
type Message struct {
	User *users.User
	Data string
}

// NewMessage ... creates a new message
func NewMessage(user *users.User, data string) *Message {
	return &Message{
		User: user,
		Data: data,
	}
}

// Receiver ... rec
type Receiver interface {
	OnMessage(message *Message)
}

// Game ... default entity to structure rooms
type Game struct {
	id    string
	title string
	rooms map[string]Room

	running    bool
	SystemUser *users.User

	OnMessageReceived chan *Message
	OnUserJoined      chan *UserJoined
	OnUserQuit        chan *UserQuit

	Receivers []Receiver

	CommandProcessor *CommandProcessor
}

var instance *Game
var once sync.Once

// Subscribe ... sub
func (game *Game) Subscribe(receiver Receiver) {
	game.Receivers = append(game.Receivers, receiver)
}

/*
func (game *Game) Unsubscribe(receiver *Receiver) {
	game.Receivers = delete(game.Receivers, receiver)
}*/

// GetInstance ... returns the usermanager instance
func GetInstance() *Game {
	once.Do(func() {
		instance = &Game{
			running:          true,
			title:            "Lair of the Dragon",
			SystemUser:       users.NewUser("LOTD", "", ""),
			CommandProcessor: NewCommandProcessor(),
			// event channels
			OnMessageReceived: make(chan *Message, 10),
			OnUserJoined:      make(chan *UserJoined, 10),
			OnUserQuit:        make(chan *UserQuit, 10),

			// game update listeners
			Receivers: make([]Receiver, 0, 10),
		}
		instance.run()
	})
	return instance
}

// CreateRoom ... processes every entity
func (game *Game) CreateRoom(title string) *Room {

	room := NewRoom()
	return room
}

// ID ... returns the id of the room
func (game *Game) ID() string {
	return game.id
}

// main game loop
func (game *Game) run() {

	go func() {
		for {
			select {
			case userJoined := <-game.OnUserJoined:

				var msg = NewMessage(game.SystemUser, userJoined.User.ID+" joined.")
				game.OnMessageReceived <- msg

			case userQuit := <-game.OnUserQuit:

				var msg = NewMessage(game.SystemUser, userQuit.User.ID+" quitted.")
				game.OnMessageReceived <- msg

			case message := <-game.OnMessageReceived:

				// only broadcast if commandprocessor didnt process it
				if !game.CommandProcessor.Process(game, message) {
					// broeadcast message
					for _, receiver := range game.Receivers {
						receiver.OnMessage(message)
					}
				}

			}
		}
	}()
}
