package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Biography string
}

func NewUser(firstName, lastName, biography string) *User {
	user := &User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Biography: biography,
	}

	return user
}
