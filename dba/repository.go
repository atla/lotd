package dba

import (
	"go.uber.org/fx"
)

// RepositoryModule ... exports dbaccess module
var RepositoryModule = fx.Provide(func(dba *DBAccess) *Repository {
	return &Repository{
		Avatars: NewAvatarDAO(dba),
		Rooms:   NewRoomDAO(dba),
		Worlds:  NewWorldDAO(dba),
	}
})

// Repository ... datatype for the ruleset
type Repository struct {
	Rooms   *RoomDAO
	Avatars *AvatarDAO
	Worlds  *WorldDAO
}
