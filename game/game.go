package game

import (
	"sync"

	"github.com/atla/lotd/users"
)

// Game ... default entity to structure rooms
type Game struct {
	id    string
	title string
	world *World

	MOTD string

	running    bool
	SystemUser *users.User

	OnMessageReceived chan *Message
	OnUserJoined      chan *UserJoined
	OnUserQuit        chan *UserQuit

	OnAvatarJoinedRoom chan *AvatarJoinedRoom
	OnAvatarLeftRoom   chan *AvatarLeftRoom

	Receivers []Receiver

	CommandProcessor *CommandProcessor

	Avatars map[string]*Avatar
}

var instance *Game
var once sync.Once

// Subscribe ... sub
func (game *Game) Subscribe(receiver Receiver) {
	game.Receivers = append(game.Receivers, receiver)
}

// Receiver ... rec
type Receiver interface {
	OnMessage(message interface{})
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
			MOTD:             "Welcome",
			SystemUser:       users.NewUser("LOTD", "", ""),
			CommandProcessor: NewCommandProcessor(),
			world:            NewWorld("LOTD"),
			// event channels
			OnMessageReceived:  make(chan *Message, 20),
			OnUserJoined:       make(chan *UserJoined, 20),
			OnUserQuit:         make(chan *UserQuit, 20),
			OnAvatarJoinedRoom: make(chan *AvatarJoinedRoom, 20),
			OnAvatarLeftRoom:   make(chan *AvatarLeftRoom, 20),

			// game update listeners
			Receivers: make([]Receiver, 0, 10),

			Avatars: make(map[string]*Avatar),
		}
		instance.run()
	})
	return instance
}

// CreateRoom ... processes every entity
func (game *Game) CreateRoom(title string) *Room {

	room := NewRoom("randomid", title, "description")
	return room
}

// ID ... returns the id of the room
func (game *Game) ID() string {
	return game.id
}

func (game *Game) sendMessage(message interface{}) {
	// broeadcast message
	for _, receiver := range game.Receivers {
		receiver.OnMessage(message)
	}
}

func (game *Game) loadAvatarForUser(user *users.User) {

	if avatar, ok := game.Avatars[user.ID]; ok {

		avatar.CurrentUser = user

		if avatar.LastKnownRoom == nil {
			avatar.LastKnownRoom = game.world.GetStartingRoom()
		}

		avatar.LastKnownRoom.Enter(avatar)

	} else {
		var newAvatar = NewAvatar()
		newAvatar.ID = user.ID
		game.Avatars[user.ID] = newAvatar

		game.loadAvatarForUser(user)
	}

}

// main game loop
func (game *Game) run() {

	go func() {
		for {
			select {
			case userJoined := <-game.OnUserJoined:

				game.loadAvatarForUser(userJoined.User)

				//TODO: remove? only inform avatars in same room?
			//	game.sendMessage(NewMessage(game.SystemUser, userJoined.User.ID+" joined."))

			case userQuit := <-game.OnUserQuit:

				_ = userQuit

			//	game.sendMessage(NewMessage(game.SystemUser, userQuit.User.ID+" quitted."))

			case avatarJoinedRoom := <-game.OnAvatarJoinedRoom:

				//	var user = avatarJoinedRoom.Avatar.CurrentUser
				//	var msg = NewMessage(nil, "=== "+avatarJoinedRoom.Room.Title+" ===\n"+avatarJoinedRoom.Room.Description)

				//	msg.ToUser = user

				game.sendMessage(avatarJoinedRoom)

			case message := <-game.OnMessageReceived:

				// only broadcast if commandprocessor didnt process it
				if !game.CommandProcessor.Process(game, message) {
					game.sendMessage(message)
				}
			}
		}
	}()
}
