package dba

import (
	"github.com/globalsign/mgo/bson"
)

// RoomDAO ... datatype for the ruleset
type RoomDAO struct {
	db         *DBAccess
	collection string
}

// NewRoomDAO ... creates a new ruleset
func NewRoomDAO(db *DBAccess) *RoomDAO {
	return &RoomDAO{
		db:         db,
		collection: "rooms",
	}
}

// FindByID ... finds a room by id
func (r *RoomDAO) FindByID(id string) (RoomDTO, error) {
	var room RoomDTO
	err := r.db.C(r.collection).FindId(bson.ObjectIdHex(id)).One(&room)
	return room, err
}

// FindAll ... find all rooms
func (r *RoomDAO) FindAll() ([]RoomDAO, error) {
	var rooms []RoomDAO
	err := r.db.C(r.collection).Find(bson.M{}).All(&rooms)
	return rooms, err
}

// Insert ... insert a room
func (r *RoomDAO) Insert(room RoomDTO) error {
	err := r.db.C(r.collection).Insert(&room)
	return err
}

// Delete ... delete a room
func (r *RoomDAO) Delete(room RoomDTO) error {
	err := r.db.C(r.collection).Remove(&room)
	return err
}

// Update ... update an existing room by id
func (r *RoomDAO) Update(room RoomDTO) error {
	err := r.db.C(r.collection).UpdateId(room.ID, &room)
	return err
}
