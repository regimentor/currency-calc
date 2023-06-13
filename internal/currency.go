package internal

import (
	"fmt"
	"log"
	"time"

	currencyapi_com "github.com/regimentor/currency-calc/internal/currencyapi.com"
)

type CurrencyId int64

type Currency struct {
	ID    CurrencyId `json:"id"`
	Slug  string     `json:"slug"`
	Value float64    `json:"value"`
	Date  time.Time  `json:"date"`
	Base  string     `json:"base"`
}

type CreateCurrencyDto struct {
	Slug  string
	Value float64
	Date  time.Time
	Base  string
}

type CurrencyStorage interface {
	GetBySlug(slug []string, date time.Time) ([]Currency, error)
	GetBySlugAndBase(slug []string, base string, date time.Time) ([]Currency, error)
	Create(currency *CreateCurrencyDto) (*Currency, error)
}

type CurrencyRepository struct {
	storage CurrencyStorage
	api     *currencyapi_com.CurrencyApiCom
}

func NewCurrencyRepository(storage CurrencyStorage, api *currencyapi_com.CurrencyApiCom) *CurrencyRepository {
	return &CurrencyRepository{storage: storage, api: api}
}

func (r *CurrencyRepository) GetBySlug(slug []string, date time.Time) ([]Currency, error) {
	log.Printf("CurrencyRepository.GetBySlug slug: %s, date: %s", slug, date)

	curs, err := r.storage.GetBySlug(slug, date)
	if err != nil {
		log.Printf("CurrencyRepository.GetBySlug, err: %v", err)

		res, err := r.api.GetCurrenciesByDate(slug, date)
		if err != nil {
			return nil, fmt.Errorf("get currency due err: %v", err)
		}

		currencies := make([]Currency, 0, len(slug))
		for _, slug := range slug {
			newCurrency := &CreateCurrencyDto{
				Slug:  slug,
				Value: res.Data[slug].Value,
				Date:  date,
				Base:  "USD",
			}

			currency, err := r.storage.Create(newCurrency)
			if err != nil {
				return nil, fmt.Errorf("get currency due err: %v", err)
			}

			currencies = append(currencies, *currency)
		}

		return currencies, nil
	}

	return curs, nil
}

func (r *CurrencyRepository) GetBySlugAndBase(slugs []string, base string, date time.Time) ([]Currency, error) {
	log.Printf("CurrencyRepository.GetBySlugAndBase slug: %s, base: %s, date: %s", slugs, base, date)

	currencies, err := r.storage.GetBySlugAndBase(slugs, base, date)
	if err != nil {
		log.Printf("CurrencyRepository.GetBySlugAndBase, err: %v", err)

		res, err := r.api.GetCurrenciesFromTo(base, slugs, date)
		if err != nil {
			return nil, fmt.Errorf("get currency due err: %v", err)
		}

		currencies := make([]Currency, 0, len(slugs))

		for _, slug := range slugs {
			newCurrency := &CreateCurrencyDto{
				Slug:  slug,
				Value: res.Data[slug].Value,
				Date:  date,
				Base:  base,
			}

			currency, err := r.storage.Create(newCurrency)
			if err != nil {
				return nil, fmt.Errorf("get currency due err: %v", err)
			}

			currencies = append(currencies, *currency)
		}

		return currencies, nil

	}

	return currencies, nil
}
