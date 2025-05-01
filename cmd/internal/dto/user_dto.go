package dto

import (
	"github.com/google/uuid"
	"github.com/mauricio-msp/go-users-api/cmd/internal/domain"
)

type UserCreateInput struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=20"`
	LastName  string `json:"last_name" validate:"required,min=2,max=20"`
	Biography string `json:"bio" validate:"required,min=20,max=450"`
}

type UserOutput struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	FullName  string    `json:"full_name"`
	Biography string    `json:"bio"`
}

func ToUser(userInput UserCreateInput) *domain.User {
	return domain.NewUser(userInput.FirstName, userInput.LastName, userInput.Biography)
}

func FromUser(user *domain.User) UserOutput {
	return UserOutput{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		FullName:  user.FirstName + " " + user.LastName,
		Biography: user.Biography,
	}
}
