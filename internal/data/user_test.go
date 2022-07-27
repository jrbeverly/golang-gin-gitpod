package data

import "testing"

func TestUserValidity(t *testing.T) {
	db := NewDatabase()
	if !IsUserExist(db, "user1", "pass1") {
		t.Fail()
	}

	if IsUserExist(db, "user2", "pass1") {
		t.Fail()
	}

	if IsUserExist(db, "user1", "") {
		t.Fail()
	}

	if IsUserExist(db, "", "pass1") {
		t.Fail()
	}

	if IsUserExist(db, "User1", "pass1") {
		t.Fail()
	}
}

func TestValidUserRegistration(t *testing.T) {
	db := NewDatabase()
	u, err := CreateUser(db, "newuser", "newpass")

	if err != nil || u.Username == "" {
		t.Fail()
	}
}

func TestInvalidUserRegistration(t *testing.T) {
	db := NewDatabase()

	u, err := CreateUser(db, "user1", "pass1")

	if err == nil || u != nil {
		t.Fail()
	}

	u, err = CreateUser(db, "newuser", "")

	if err == nil || u != nil {
		t.Fail()
	}
}

func TestUsernameAvailability(t *testing.T) {
	db := NewDatabase()

	if !isUsernameAvailable(db, "newuser") {
		t.Fail()
	}

	if isUsernameAvailable(db, "user1") {
		t.Fail()
	}

	CreateUser(db, "newuser", "newpass")
	if isUsernameAvailable(db, "newuser") {
		t.Fail()
	}
}
