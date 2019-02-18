package dba

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// RoomDTO ... default entity to structure rooms
type RoomDTO struct {
	ID          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	DateCreated time.Time     `json:"dateCreated,omitempty"`
}

// NewRoomDTO ... creates and returns a new room instance
func NewRoomDTO(id string, title string, description string) *RoomDTO {
	return &RoomDTO{
		ID:          bson.NewObjectId(),
		Title:       title,
		Description: description,
		DateCreated: time.Now(),
	}
}
