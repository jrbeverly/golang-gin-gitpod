package v1

// Represents an article published to the website
type Article struct {
	// Unique identifier for the article
	ID int `json:"id"`

	// Tile of the article
	Title string `json:"title"`

	// The HTML body of the article
	Content string `json:"content"`
}
