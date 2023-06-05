package internal

type User struct {
	ID     uint
	ApiKey string
}

type UserStorage interface {
	GetById(id uint) (User, error)
	GetByApiKey(apiKey string) (User, error)
	Create(u *User) (User, error)
}

type UserRepository struct {
	storage UserStorage
}

func NewUserRepository(storage UserStorage) *UserRepository {
	return &UserRepository{storage: storage}
}

func (r *UserRepository) Create(u *User) (*User, error) {

	return u, nil
}
