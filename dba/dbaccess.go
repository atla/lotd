package dba

import (
	"go.uber.org/fx"

	"github.com/globalsign/mgo"
)

const (
	databaseName = "lotd"
)

// AccessModule ... exports dbaccess module
var AccessModule = fx.Provide(func() (*DBAccess, error) {

	if session, err := mgo.Dial("0.0.0.0:27017"); err == nil {
		return NewDBAccess(session, databaseName), nil
	} else {
		return nil, err
	}

})

// DBAccess ... datatype for the ruleset
type DBAccess struct {
	Session  *mgo.Session
	Database string
}

// NewDBAccess ... creates a new ruleset
func NewDBAccess(session *mgo.Session, dbName string) *DBAccess {
	return &DBAccess{
		Session:  session,
		Database: dbName,
	}
}

// DB ... returns database instance
func (dba *DBAccess) DB() *mgo.Database {
	return dba.Session.DB(dba.Database)
}

// C ... returns collection
func (dba *DBAccess) C(collection string) *mgo.Collection {
	return dba.DB().C(collection)

}
