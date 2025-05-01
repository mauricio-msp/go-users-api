package domain

type UserRepository interface {
	Create(user *User) error
	FindAll() ([]*User, error)
	FindByID(id string) (*User, error)
	Update(id string, user *User) error
	Delete(id string) error
}
