package data

import (
	"errors"
	"strings"

	v1 "github.com/jrbeverly/cobra-cmd-with-docs/internal/models/v1"
)

// Register a new user with the given username and password
func CreateUser(db *InMemoryDatabase, username string, password string) (*v1.User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	}

	if !isUsernameAvailable(db, username) {
		return nil, errors.New("The username isn't available")
	}

	u := v1.User{Username: username, Password: password}

	db.Users = append(db.Users, u)

	return &u, nil
}

// Check if the supplied username is available
func isUsernameAvailable(db *InMemoryDatabase, username string) bool {
	for _, u := range db.Users {
		if u.Username == username {
			return false
		}
	}
	return true
}

// Check if the username and password combination is valid
func IsUserExist(db *InMemoryDatabase, username, password string) bool {
	for _, u := range db.Users {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}
