package dba

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// WorldDTO ... default entity to structure rooms
// Everything regarding content and live/dynamic data such as items, avatars, room shall be
// managed from the World class - all generic game/message/command related things will reside in the game
// class
type WorldDTO struct {
	ID          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	DateCreated time.Time     `json:"dateCreated,omitempty"`
	Description string        `json:"description,omitempty"`

	Rooms          []*RoomDTO   `json:"rooms,omitempty"`
	Avatars        []*AvatarDTO `json:"avatars,omitempty"`
	StartingRoomID string       `json:"startingRoomId,omitempty"`
}

// NewWorldDTO ... creates and returns a new room instance
func NewWorldDTO(description string) *WorldDTO {
	return &WorldDTO{
		ID:          bson.NewObjectId(),
		Description: description,
		DateCreated: time.Now(),
	}
}
