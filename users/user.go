package users

import (
	"fmt"
)

// User ... global user struct to control logins
type User struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Active   bool   `json:"active"`
}

// NewUser ... creates a new user
func NewUser(id string, password string, email string) *User {
	return &User{
		ID:       id,
		Password: password,
		Email:    email,
		Active:   false,
	}
}

// Login ... login
func (user *User) Login() {
	fmt.Println("User logged in: " + user.ID)
	user.Active = true
}

// Logout ... logout
func (user *User) Logout() {
	fmt.Println("User logged out: " + user.ID)
	user.Active = false
}

// Timeout ... timeout
func (user *User) Timeout() {
	fmt.Println("User timed out: " + user.ID)
	user.Active = false
}
