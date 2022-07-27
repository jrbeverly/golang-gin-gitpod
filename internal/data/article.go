package data

import (
	"errors"

	v1 "github.com/jrbeverly/cobra-cmd-with-docs/internal/models/v1"
)

// Return a list of all the articles
func ListArticles(db *InMemoryDatabase) []v1.Article {
	return db.Articles
}

// Fetch an article based on the ID supplied
func GetArticleByID(db *InMemoryDatabase, id int) (*v1.Article, error) {
	for _, a := range db.Articles {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

// Create a new article with the title and content provided
func NewArticle(db *InMemoryDatabase, title, content string) (*v1.Article, error) {
	a := v1.Article{ID: len(db.Articles) + 1, Title: title, Content: content}
	db.Articles = append(db.Articles, a)
	return &a, nil
}
