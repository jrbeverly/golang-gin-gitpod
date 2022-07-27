package data

import (
	v1 "github.com/jrbeverly/cobra-cmd-with-docs/internal/models/v1"
)

// Represents an in-memory struct acting as a stub database
type InMemoryDatabase struct {
	// The articles contained within the database
	Articles []v1.Article

	// The users contained within the database
	Users []v1.User
}

// Fetch an article based on the ID supplied
func NewDatabase() *InMemoryDatabase {
	database := &InMemoryDatabase{}
	database.Articles = []v1.Article{
		{ID: 1, Title: "Article 1", Content: "Article 1 body"},
		{ID: 2, Title: "Article 2", Content: "Article 2 body"},
	}

	database.Users = []v1.User{
		{Username: "admin", Password: "admin"},
		{Username: "user1", Password: "pass1"},
		{Username: "user2", Password: "pass2"},
		{Username: "user3", Password: "pass3"},
	}

	return database
}
