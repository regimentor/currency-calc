package postgresql

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/regimentor/currency-calc/internal"
)

type UserStorage struct {
	connection *pgxpool.Pool
}

func NewUserStorage(connection *pgxpool.Pool) *UserStorage {
	return &UserStorage{connection: connection}
}

func (s *UserStorage) GetById(id uint) (internal.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *UserStorage) GetByApiKey(apiKey string) (internal.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *UserStorage) Create(u *internal.User) (internal.User, error) {
	//TODO implement me
	panic("implement me")
}
