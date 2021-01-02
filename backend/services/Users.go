package services

import (
	"fmt"

	"github.com/conduit-golang/backend/model"
)

// CreateUser add user in users table
func CreateUser(user model.User) {
	fmt.Println(user)
}
