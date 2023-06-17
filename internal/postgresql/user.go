package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/regimentor/currency-calc/internal"
	"log"
)

type UserStorage struct {
	connection *pgxpool.Pool
}

func NewUserStorage(connection *pgxpool.Pool) *UserStorage {
	return &UserStorage{connection: connection}
}

func (s *UserStorage) GetById(ctx context.Context, id internal.UserId) (*internal.User, error) {
	log.Printf("UserStorage.GetById: %v", id)

	query := `
		select id, api_key from users where id = $1;
	`

	user := &internal.User{}

	row := s.connection.QueryRow(ctx, query, id)
	if err := row.Scan(&user.ID, &user.ApiKey); err != nil {
		return nil, fmt.Errorf("get user due err: %v", err)
	}

	log.Printf("UserStorage.GetById, got user: %v", user)

	return user, nil
}

func (s *UserStorage) GetByApiKey(ctx context.Context, apiKey internal.ApiKey) (*internal.User, error) {
	log.Printf("UserStorage.GetByApiKey: %v", apiKey)

	query := `
		select id, api_key from users where api_key = $1;
	`

	user := &internal.User{}

	row := s.connection.QueryRow(ctx, query, apiKey)
	if err := row.Scan(&user.ID, &user.ApiKey); err != nil {
		return nil, fmt.Errorf("get user due err: %v", err)
	}

	log.Printf("UserStorage.GetByApiKey, got user: %v", user)

	return user, nil
}

func (s *UserStorage) Create(ctx context.Context, u internal.CreateUserDto) (*internal.User, error) {
	log.Printf("UserStorage.Create: %v", u)

	createUserQuery := `
		insert into users (api_key) values ($1) returning id, api_key;
	`
	newUser := &internal.User{}

	row := s.connection.QueryRow(ctx, createUserQuery, u.ApiKey)

	if err := row.Scan(&newUser.ID, &newUser.ApiKey); err != nil {
		return nil, fmt.Errorf("create user due err: %v", err)
	}

	log.Printf("UserStorage.Create, created user: %v", u)
	return newUser, nil
}
