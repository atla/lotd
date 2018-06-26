package game

import (
	"github.com/atla/lotd/users"
)

// Avatar ... default entity to structure rooms
type Avatar struct {
	ID            string `json:"id"`
	LastKnownRoom *Room
	CurrentUser   *users.User
}

// NewAvatar ... creates and returns a new room instance
func NewAvatar() *Avatar {
	return &Avatar{
		LastKnownRoom: nil,
	}
}
