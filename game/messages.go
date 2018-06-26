package game

import "github.com/atla/lotd/users"

//UserJoined ... player joined event
type UserJoined struct {
	User *users.User
}

//UserQuit ... player joined event
type UserQuit struct {
	User *users.User
}

// AvatarJoinedRoom ... asdd
type AvatarJoinedRoom struct {
	Avatar *Avatar
	Room   *Room
}

// AvatarLeftRoom ... asdd
type AvatarLeftRoom struct {
	Avatar *Avatar
	Room   *Room
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
	FromUser *users.User
	ToUser   *users.User
	Data     string
}

// NewMessage ... creates a new message
func NewMessage(fromUser *users.User, data string) *Message {
	return &Message{
		FromUser: fromUser,
		Data:     data,
	}
}
