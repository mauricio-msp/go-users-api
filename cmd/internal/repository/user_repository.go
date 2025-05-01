package repository

import (
	"fmt"
	"sync"

	"github.com/mauricio-msp/go-users-api/cmd/internal/domain"
)

type UserRepository struct {
	db  map[string]*domain.User
	mux sync.RWMutex
}

func NewUserRepository(db map[string]*domain.User) *UserRepository {
	return &UserRepository{db: db}
}

// func (ur *UserRepository) PrintAll() {
// 	ur.mux.RLock()
// 	defer ur.mux.RUnlock()

// 	fmt.Println("---- Users in Memory ----")
// 	copy := make(map[string]*domain.User)
// 	maps.Copy(copy, ur.db)

// 	for k, v := range copy {
// 		fmt.Printf("ID: %s | Nome: %s %s | Bio: %s\n", k, v.FirstName, v.LastName, v.Biography)
// 	}
// }

func (ur *UserRepository) Create(user *domain.User) error {
	ur.mux.Lock()
	defer ur.mux.Unlock()

	if _, exists := ur.db[user.ID.String()]; exists {
		return fmt.Errorf("user with ID %s already exists", user.ID)
	}

	ur.db[user.ID.String()] = user

	return nil
}

func (ur *UserRepository) FindAll() ([]*domain.User, error) {
	ur.mux.RLock()
	defer ur.mux.RUnlock()

	var users []*domain.User
	for _, user := range ur.db {
		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) FindByID(id string) (*domain.User, error) {
	ur.mux.RLock()
	defer ur.mux.RUnlock()

	user, exists := ur.db[id]
	if !exists {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil
}

func (ur *UserRepository) Update(id string, user *domain.User) error {
	ur.mux.Lock()
	defer ur.mux.Unlock()

	if _, exists := ur.db[id]; !exists {
		return fmt.Errorf("user with ID %s not found", id)
	}

	ur.db[id] = user
	return nil
}

func (ur *UserRepository) Delete(id string) error {
	ur.mux.Lock()
	defer ur.mux.Unlock()

	if _, exists := ur.db[id]; !exists {
		return fmt.Errorf("user with ID %s not found", id)
	}

	delete(ur.db, id)
	return nil
}
