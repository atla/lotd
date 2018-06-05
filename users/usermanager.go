package users

import (
	"strings"
	"sync"
)

// UserManager ... global user struct to control logins
type UserManager struct {
	users []*User
}

var instance *UserManager
var once sync.Once

// GetInstance ... returns the usermanager instance
func GetInstance() *UserManager {
	once.Do(func() {
		instance = &UserManager{}

		instance.setupTestUsers()
	})
	return instance
}

func (userManager *UserManager) setupTestUsers() {

	userManager.AddUser(NewUser("atla", "test", "koerner.marcus@gmail.com"))
	userManager.AddUser(NewUser("parci", "test", "parci.val@gmail.com"))
	userManager.AddUser(NewUser("bilbo", "test", "bil.bo@gmail.com"))

}

// AddUser .. adds a user
func (userManager *UserManager) AddUser(user *User) {
	userManager.users = append(userManager.users, user)
}

// GetAllUsers ... asd
func (userManager *UserManager) GetAllUsers(id string) []*User {
	return userManager.users
}

// FindUserByID ... finds a user by id
func (userManager *UserManager) FindUserByID(id string) *User {

	for _, user := range userManager.users {
		if strings.HasPrefix(id, user.ID) {
			return user
		}

	}

	return nil
}
