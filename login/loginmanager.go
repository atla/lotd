package login

import (
	"log"
)

// LoginManager ... asd
type LoginManager struct {
}

// NewLoginManager ... creates new login manager
func NewLoginManager() *LoginManager {
	return &LoginManager{}
}

// Login ... uses a baseauth with userid and password hash
func (auth *LoginManager) Login(user string, password string) bool {

	log.Println("Received login challenge with user " + user + " and password " + password)

	

	return true

}
