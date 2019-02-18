package dba

import (
	"github.com/globalsign/mgo/bson"
)

// AvatarDAO ... datatype for the ruleset
type AvatarDAO struct {
	db         *DBAccess
	collection string
}

// NewAvatarDAO ... creates a new ruleset
func NewAvatarDAO(db *DBAccess) *AvatarDAO {
	return &AvatarDAO{
		db:         db,
		collection: "avatars",
	}
}

// FindByID ... finds a room by id
func (a *AvatarDAO) FindByID(id string) (AvatarDTO, error) {
	var avatar AvatarDTO
	err := a.db.C(a.collection).FindId(bson.ObjectIdHex(id)).One(&avatar)
	return avatar, err
}

// FindAll ... find all rooms
func (a *AvatarDAO) FindAll() ([]AvatarDAO, error) {
	var avatars []AvatarDAO
	err := a.db.C(a.collection).Find(bson.M{}).All(&avatars)
	return avatars, err
}

// Insert an AvatarDTO into the db
func (a *AvatarDAO) Insert(avatar AvatarDTO) error {
	err := a.db.C(a.collection).Insert(&avatar)
	return err
}

// Delete an avatar
func (a *AvatarDAO) Delete(avatar AvatarDTO) error {
	err := a.db.C(a.collection).Remove(&avatar)
	return err
}

// Update ... update an existing room by id
func (a *AvatarDAO) Update(avatar AvatarDTO) error {
	err := a.db.C(a.collection).UpdateId(avatar.ID, &avatar)
	return err
}

// InsertOrUpdate ... insert or update an existing room by id
func (a *AvatarDAO) InsertOrUpdate(avatar AvatarDTO) error {

	err := a.Update(avatar)

	if err != nil {
		return a.Insert(avatar)
	}

	return err
}
