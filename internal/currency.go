package internal

import (
	"fmt"
	"time"
)

type Currency struct {
}

type CreateCurrencyDto struct {
}

type CurrencyStorage interface {
	GetByDate(currencies []string, date time.Time) (*Currency, error)
	GetFromTo(from, to string) ([]Currency, error)
	Create(currency *CreateCurrencyDto) (*Currency, error)
}

type CurrencyRepository struct {
	storage CurrencyStorage
}

func NewCurrencyRepository(storage CurrencyStorage) *CurrencyRepository {
	return &CurrencyRepository{storage: storage}
}

func (r *CurrencyRepository) GetByDate(currencies []string, date time.Time) (*Currency, error) {
	currency, err := r.storage.GetByDate(currencies, date)
	if err != nil {
		return nil, fmt.Errorf("create user from storage due err: %v", err)
	}

	return currency, nil
}
