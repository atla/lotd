package dba

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// AvatarDTO ... dto to store data related to avatars
// Avatars can be either controlled by Players/Users or be attached/belong to bots
// Once a user is logged in he automatically gets attached his last used aavatar
type AvatarDTO struct {
	ID bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`

	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CreationDate time.Time `json:"creationDate,omitempty"`

	CurrentRoomID string `json:"lastKnownRoomID,omitempty" bson:"lastKnownRoomID,omitempty"`
	UserID        string `json:"userID"`
}

// NewAvatarDTO ... creates and returns a new room instance
func NewAvatarDTO(name string, description string, userID string) *AvatarDTO {
	return &AvatarDTO{
		ID:          bson.NewObjectId(),
		Name:        name,
		Description: description,
		UserID:      userID,
	}
}
