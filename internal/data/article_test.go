package data

import "testing"

func TestGetArticleByID(t *testing.T) {
	db := NewDatabase()

	expected, err := NewArticle(db, "New test title", "New test content")

	actual, err := GetArticleByID(db, expected.ID)

	if err != nil || expected.ID != actual.ID || expected.Title != actual.Title || expected.Content != actual.Content {
		t.Fail()
	}
}

func TestCreateNewArticle(t *testing.T) {
	db := NewDatabase()

	originalLength := len(ListArticles(db))
	a, err := NewArticle(db, "New test title", "New test content")

	allArticles := ListArticles(db)
	newLength := len(allArticles)

	if err != nil || newLength != originalLength+1 ||
		a.Title != "New test title" || a.Content != "New test content" {

		t.Fail()
	}
}
