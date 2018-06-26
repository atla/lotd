package users

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/atla/lotd/ldb"
)

// UserManager ... global user struct to control logins
type UserManager struct {
	users    []User
	DBHelper *ldb.Helper
}

var instance *UserManager
var once sync.Once

// GetInstance ... returns the usermanager instance
func GetInstance() *UserManager {
	once.Do(func() {
		instance = &UserManager{}

		instance.DBHelper = ldb.NewHelper("db/users")

		instance.setupTestUsers()
	})
	return instance
}

func (userManager *UserManager) setupTestUsers() {

	userManager.AddUser(NewUser("atla", "test", "koerner.marcus@gmail.com"))
	userManager.AddUser(NewUser("parci", "test", "parci.val@gmail.com"))
	userManager.AddUser(NewUser("bilbo", "test", "bil.bo@gmail.com"))
	userManager.AddUser(NewUser("guest", "guest", "guest@random.com"))
}

// AddUser .. adds a user
func (userManager *UserManager) AddUser(user *User) {
	//userManager.users = append(userManager.users, user)

	ub, err := json.Marshal(*user)

	log.Println("ADDUSER: " + string(ub))

	if err == nil {
		userManager.DBHelper.Put(user.ID, ub)
	} else {
		log.Println("Error marshalling user " + err.Error())
	}
}

// GetAllActiveUsers ... asd
func (userManager *UserManager) GetAllActiveUsers(id string) []User {
	return userManager.users
}

// FindUserByID ... finds a user by id
func (userManager *UserManager) FindUserByID(id string) (*User, error) {

	userData, err := userManager.DBHelper.Get(id)

	if err != nil {
		return nil, err
	}

	var user User
	err2 := json.Unmarshal(userData, user)

	if err2 != nil {
		log.Println("Error unmarshalling user " + err2.Error())
		return nil, err2
	}

	return &user, nil
}
