package dba

import (
	"github.com/globalsign/mgo/bson"
)

// WorldDAO ... datatype for the ruleset
type WorldDAO struct {
	db         *DBAccess
	collection string
}

// NewWorldDAO ... creates a new ruleset
func NewWorldDAO(db *DBAccess) *WorldDAO {
	return &WorldDAO{
		db:         db,
		collection: "worlds",
	}
}

// FindByID ... finds a room by id
func (r *WorldDAO) FindByID(id string) (WorldDTO, error) {
	var world WorldDTO
	err := r.db.C(r.collection).FindId(bson.ObjectIdHex(id)).One(&world)
	return world, err
}

// Insert ... insert a room
func (r *WorldDAO) Insert(world WorldDTO) error {
	err := r.db.C(r.collection).Insert(&world)
	return err
}

// Delete deletes a world
func (r *WorldDAO) Delete(world WorldDTO) error {
	err := r.db.C(r.collection).Remove(&world)
	return err
}

// Update update an existing world by id
func (r *WorldDAO) Update(world WorldDTO) error {
	err := r.db.C(r.collection).UpdateId(world.ID, &world)
	return err
}
