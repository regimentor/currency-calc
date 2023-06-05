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

func (s *UserStorage) GetById(id internal.UserId) (internal.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *UserStorage) GetByApiKey(apiKey internal.ApiKey) (internal.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *UserStorage) Create(u internal.CreateUserDto) (internal.User, error) {
	//TODO implement me
	panic("implement me")
}
