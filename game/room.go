package game

import (
	"fmt"

	"github.com/segmentio/ksuid"
)

// Room ... default entity to structure rooms
type Room struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// NewRoom ... creates and returns a new room instance
func NewRoom() *Room {
	return &Room{
		ID: ksuid.New().String(),
	}
}

// Enter ... processes every entity
func (room *Room) Enter() {
	fmt.Println("Someone entered " + room.ID)
}
