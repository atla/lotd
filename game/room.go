package game

import (
	"fmt"
)

// Room ... default entity to structure rooms
type Room struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// NewRoom ... creates and returns a new room instance
func NewRoom() *Room {
	return &Room{}
}

// Enter ... processes every entity
func (room *Room) Enter() {
	fmt.Println("Someone entered " + room.ID)
}
