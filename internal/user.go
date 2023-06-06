package internal

import "fmt"

type UserId uint
type ApiKey string

func (r *ApiKey) generate() {
	panic("TODO: implement me")
}

type User struct {
	ID     UserId
	ApiKey ApiKey
}

type CreateUserDto struct {
}

type UserStorage interface {
	GetById(id UserId) (*User, error)
	GetByApiKey(apiKey ApiKey) (*User, error)
	Create(u CreateUserDto) (*User, error)
}

type UserRepository struct {
	storage UserStorage
}

func NewUserRepository(storage UserStorage) *UserRepository {
	return &UserRepository{storage: storage}
}

func (r *UserRepository) Create(u CreateUserDto) (*User, error) {
	newUser, err := r.storage.Create(u)
	if err != nil {
		return nil, fmt.Errorf("create user from storage due err: %v", err)
	}

	return newUser, nil
}
