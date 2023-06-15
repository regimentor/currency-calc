package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func NewConnection(ctx context.Context, user, password, db, host string) (*pgxpool.Pool, error) {
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s", user, password, host, db)
	log.Printf("connection to database with url: %v", dbUrl)
	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		return nil, fmt.Errorf("create connection to database due err: %v", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping connection to database due err: %v", err)
	}

	return pool, nil
}
