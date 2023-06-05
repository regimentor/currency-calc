package internal

type UserId uint
type ApiKey string

type User struct {
	ID     UserId
	ApiKey ApiKey
}

type CreateUserDto struct {
}

type UserStorage interface {
	GetById(id UserId) (User, error)
	GetByApiKey(apiKey ApiKey) (User, error)
	// TODO: CreateUserDto
	Create(u CreateUserDto) (User, error)
}

type UserRepository struct {
	storage UserStorage
}

func NewUserRepository(storage UserStorage) *UserRepository {
	return &UserRepository{storage: storage}
}

func (r *UserRepository) Create(u CreateUserDto) (*User, error) {

	return nil, nil
}
