package v1

// Represents a user within the blog
type User struct {
	// The unique name of the user
	Username string `json:"username"`

	// The clear-text password for the service
	Password string `json:"-"`
}
