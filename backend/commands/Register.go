package commands

// Register /register user request
type Register struct {
	User user
}

type user struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}