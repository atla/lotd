package game

import (
	"github.com/atla/lotd/dba"
	"github.com/atla/lotd/users"
)

// Avatar ... default active entity that moves in the world
// Avatars can be either controlled by Players/Users or be attached/belong to bots
// Once a user is logged in he automatically gets attached his last used aavatar
type Avatar struct {
	data        *dba.AvatarDTO
	avatarDAO   *dba.AvatarDAO
	CurrentUser *users.User `json:"id,omitempty" bson:"-"`
}

// NewAvatar ... creates and returns a new room instance
func NewAvatar() *Avatar {
	return &Avatar{
		data:        nil,
		avatarDAO:   nil,
		CurrentUser: nil,
	}
}

// onChange ... updates data in the database
func (avatar *Avatar) onChange() err {

	if avatar.avatarDAO != nil {
		return avatar.avatarDAO.InsertOrUpdate(*avatar.data)
	}
}

// Description ... returns the description
func (avatar *Avatar) Description() string {
	return avatar.data.Description
}

// SetCurrentRoomID ... returns the name
func (avatar *Avatar) SetCurrentRoomID(currentRoomID string) err {

	avatar.data.CurrentRoomID = currentRoomID
	return avatar.onChange()
}

// Name ... returns the name
func (avatar *Avatar) Name() string {
	return avatar.data.Name
}

// LoadAvatar ... loads and returns the avatar
func LoadAvatar(avatars *dba.AvatarDAO, id string) (*Avatar, error) {

	data, err := avatars.FindByID(id)

	if err != nil {
		return nil, err
	}

	return &Avatar{
		data:        &data,
		avatarDAO:   avatars,
		CurrentUser: nil,
	}, nil
}
