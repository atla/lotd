package game

import (
	"encoding/json"

	"github.com/atla/lotd/ldb"
)

// World ... default entity to structure rooms
// Everything regarding content and live/dynamic data such as items, avatars, room shall be
// managed from the World class - all generic game/message/command related things will reside in the game
// class
type World struct {
	ID string

	Rooms          map[string]*Room   `json:rooms"`
	Avatars        map[string]*Avatar `json:avatars"`
	StartingRoomID string

	RoomDB   *ldb.Helper
	AvatarDB *ldb.Helper
}

func createStartingRoom() *Room {

	return NewRoom("genesis", "Entry into the world",
		`You find yourself on a bright wide open clearing. Trees behind you and a stoney road ahead,
in the distance you can see the bright lights of a village.
		
From here you can either head (n)orth or (s)outh.`)

}

// restore from database files
func (world *World) restoreFromDatabase() {
	world.RoomDB = ldb.NewHelper(world.ID + "/rooms/")
	world.AvatarDB = ldb.NewHelper(world.ID + "/avatars/")
}

// GetRoom ...
func (world *World) GetRoom(id string) (Room, error) {

	var room Room

	if world.Rooms[id] != nil {
		return *world.Rooms[id], nil
	}

	roomData, err := world.RoomDB.Get(id)

	if err != nil {
		return room, err
	}

	err2 := json.Unmarshal(roomData, room)

	if err2 != nil {
		world.Rooms[id] = &room
	}

	return room, err2
}

// NewWorld ... creates and returns a new room instance
func NewWorld(id string) *World {
	world := &World{
		ID:    id,
		Rooms: make(map[string]*Room),
	}

	world.restoreFromDatabase()

	startingRoom := createStartingRoom()
	world.AddRoom(*startingRoom)
	world.StartingRoomID = startingRoom.ID

	return world
}

// AddRoom ... adds a room to the world
func (world *World) AddRoom(room Room) {

}

// GetStartingRoom ... returns the StartingRoom for this world
func (world *World) GetStartingRoom() *Room {
	return world.GetStartingRoom()
}
