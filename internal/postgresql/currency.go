package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/regimentor/currency-calc/internal"
	"log"
	"time"
)

type CurrencyStorage struct {
	connection *pgxpool.Pool
}

func NewCurrencyStorage(connection *pgxpool.Pool) *CurrencyStorage {
	return &CurrencyStorage{connection: connection}
}

func (c *CurrencyStorage) GetBySlug(slug []string, date time.Time) ([]internal.Currency, error) {
	log.Printf("CurrencyStorage.GetByDate slug: %s, base: %s", slug, date)

	year, month, day := date.Date()
	dateStr := fmt.Sprintf("%d-%d-%d", year, month, day)

	query := `
		select id, slug, value, date, base from currencies 
		where slug = any($1) and date = $2;
	`

	currencies := make([]internal.Currency, 0, len(slug))

	rows, err := c.connection.Query(context.Background(), query, slug, dateStr)
	if err != nil {
		return nil, fmt.Errorf("get currency due err: %v", err)
	}

	for rows.Next() {
		currency := internal.Currency{}
		err := rows.Scan(&currency.ID, &currency.Slug, &currency.Value, &currency.Date, &currency.Base)
		if err != nil {
			return nil, fmt.Errorf("get currency due err: %v", err)
		}

		currencies = append(currencies, currency)
	}

	if len(currencies) == 0 {
		return nil, fmt.Errorf("get currency due err: not found")
	}

	log.Printf("CurrencyStorage.GetByDate, got currencies: %v", currencies)

	return currencies, nil
}

func (c *CurrencyStorage) GetBySlugAndBase(slug, base string, date time.Time) (*internal.Currency, error) {
	log.Printf("CurrencyStorage.GetBySlugAndBase slug: %s, base: %s, date: %s", slug, base, date)

	year, month, day := date.Date()
	dateStr := fmt.Sprintf("%d-%d-%d", year, month, day)

	query := `
		select id, slug, value, date, base from currencies 
		where slug = $1 and base = $2 and date = $3;
	`

	currency := &internal.Currency{}
	row := c.connection.QueryRow(context.Background(), query, slug, base, dateStr)

	if err := row.Scan(&currency.ID, &currency.Slug, &currency.Value, &currency.Date, &currency.Base); err != nil {
		return nil, fmt.Errorf("get currency due err: %v", err)
	}

	log.Printf("CurrencyStorage.GetBySlugAndBase, got currency: %v", currency)

	return currency, nil
}

func (c *CurrencyStorage) Create(currency *internal.CreateCurrencyDto) (*internal.Currency, error) {
	log.Printf("CurrencyStorage.Create currency: %v", currency)

	query := `
		insert into currencies (slug, value, date, base) 
		values ($1, $2, $3, $4) returning id, slug, value, date, base;
	`

	newCurrency := &internal.Currency{}

	row := c.connection.QueryRow(context.Background(), query, currency.Slug, currency.Value, currency.Date, currency.Base)
	if err := row.Scan(&newCurrency.ID, &newCurrency.Slug, &newCurrency.Value, &newCurrency.Date, &newCurrency.Base); err != nil {
		return nil, fmt.Errorf("create currency due err: %v", err)
	}

	log.Printf("CurrencyStorage.Create, got currency: %v", newCurrency)
	return newCurrency, nil
}
