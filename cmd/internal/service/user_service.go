package service

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/mauricio-msp/go-users-api/cmd/internal/domain"
	"github.com/mauricio-msp/go-users-api/cmd/internal/dto"
)

type UserService struct {
	repository domain.UserRepository
	validator  *validator.Validate
}

func NewUserService(repository domain.UserRepository) *UserService {
	return &UserService{
		repository: repository,
		validator:  validator.New(),
	}
}

func (us *UserService) CreateUser(input dto.UserCreateInput) error {
	if err := us.validator.Struct(input); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	userInput := dto.ToUser(input)

	err := us.repository.Create(userInput)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (us *UserService) FindAll() ([]*dto.UserOutput, error) {
	users, err := us.repository.FindAll()
	if err != nil {
		return nil, err
	}

	usersOutput := make([]*dto.UserOutput, 0, len(users))
	for _, user := range users {
		userOutput := dto.FromUser(user)
		usersOutput = append(usersOutput, &userOutput)
	}

	return usersOutput, nil
}

func (us *UserService) FindByID(id string) (*dto.UserOutput, error) {
	userExists, err := us.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	userOutput := dto.FromUser(userExists)

	return &userOutput, nil
}

func (us *UserService) Delete(id string) error {
	err := us.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) Update(id string, input dto.UserCreateInput) (*dto.UserOutput, error) {
	if err := us.validator.Struct(input); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	userInput := dto.ToUser(input)

	err := us.repository.Update(id, userInput)
	if err != nil {
		return nil, err
	}

	userOutput := dto.FromUser(userInput)

	return &userOutput, nil
}
