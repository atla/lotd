package login

import (
	"log"

	"github.com/atla/lotd/users"
)

// LoginManager ... asd
type LoginManager struct {
	userManager *users.UserManager
}

// NewLoginManager ... creates new login manager
func NewLoginManager() *LoginManager {
	return &LoginManager{
		userManager: users.GetInstance(),
	}
}

// Login ... uses a baseauth with userid and password hash
func (auth *LoginManager) Login(user string, password string) bool {

	log.Println("Received login challenge with user " + user + " and password " + password)

	u, err := auth.userManager.FindUserByID(user)

	if err != nil {
		log.Println("No such user " + user)
		return false
	}

	u.Active = true

	return true

}
