package postgresql

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/regimentor/currency-calc/internal"
	"time"
)

type CurrencyStorage struct {
	connection *pgxpool.Pool
}

func (c *CurrencyStorage) GetByDate(currencies []string, date time.Time) (internal.Currency, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CurrencyStorage) GetFromTo(from, to string) ([]internal.Currency, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CurrencyStorage) Create(currency *internal.CreateCurrencyDto) (*internal.Currency, error) {
	//TODO implement me
	panic("implement me")
}
