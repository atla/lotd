package game

import (
	"fmt"

	"github.com/segmentio/ksuid"
)

// Room ... default entity to structure rooms
type Room struct {
	id          string `json:"id"`
	title       string `json:"title"`
	description string `json:"description"`
}

// NewRoom ... creates and returns a new room instance
func NewRoom() *Room {
	return &Room{
		id: ksuid.New().String(),
	}
}

// Process ... processes every entity
func (room *Room) Enter() {
	fmt.Println("Someone entered " + room.id)
}

// ID ... returns the id of the room
func (room *Room) ID() string {
	return room.id
}

// Description ... returns the description of the room
func (room *Room) Description() string {
	return room.description
}

// Title ... returns the title of the room
func (room *Room) Title() string {
	return room.title
}
