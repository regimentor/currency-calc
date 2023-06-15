package internal

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

type UserId uint
type ApiKey string

func (r *ApiKey) Generate() {
	timestamp := time.Now().Unix()

	byteArray := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		byteArray[i] = byte(timestamp & 0xff)
		timestamp >>= 8
	}

	// Вычисление хэша SHA256
	hash := sha256.Sum256(byteArray)

	hex.EncodeToString(hash[:])

	*r = ApiKey(hex.EncodeToString(hash[:]))
}

type User struct {
	ID     UserId `json:"id"`
	ApiKey ApiKey `json:"apiKey"`
}

type CreateUserDto struct {
	ApiKey ApiKey
}

type UserStorage interface {
	GetById(ctx context.Context, id UserId) (*User, error)
	GetByApiKey(ctx context.Context, apiKey ApiKey) (*User, error)
	Create(ctx context.Context, u CreateUserDto) (*User, error)
}

type UserRepository struct {
	storage UserStorage
}

func NewUserRepository(storage UserStorage) *UserRepository {
	return &UserRepository{storage: storage}
}

func (r *UserRepository) Create(ctx context.Context, u CreateUserDto) (*User, error) {
	log.Printf("UserRepository:create user %s", u.ApiKey)
	newUser, err := r.storage.Create(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("create user from storage due err: %v", err)
	}

	return newUser, nil
}

func (r *UserRepository) GetById(ctx context.Context, id UserId) (*User, error) {
	log.Printf("UserRepository:GetById %d", id)
	user, err := r.storage.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get user from storage due err: %v", err)
	}

	return user, nil
}

func (r *UserRepository) GetByApiKey(ctx context.Context, apiKey ApiKey) (*User, error) {
	log.Printf("UserRepository:GetByApiKey %s", apiKey)
	user, err := r.storage.GetByApiKey(ctx, apiKey)
	if err != nil {
		return nil, fmt.Errorf("get user from storage due err: %v", err)
	}

	return user, nil
}
