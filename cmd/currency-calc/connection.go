package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func NewConnection(user, password, db string) (*pgxpool.Pool, error) {
	dbUrl := fmt.Sprintf("postgresql://%s:%s@localhost:5432/%s", user, password, db)
	log.Printf("connection to database with url: %v", dbUrl)
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, fmt.Errorf("create connection to database due err: %v", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("ping connection to database due err: %v", err)
	}

	return pool, nil
}
