package internal

import (
	"context"
	"fmt"
	"github.com/regimentor/currency-calc/internal/models"
	"log"
	"time"

	"github.com/regimentor/currency-calc/internal/currencyapi.com"
)

type CurrencyStorage interface {
	GetBySlug(ctx context.Context, slug []string, date time.Time) ([]models.Currency, error)
	GetBySlugAndBase(ctx context.Context, slug []string, base string, date time.Time) ([]models.Currency, error)
	Create(ctx context.Context, currency *models.CreateCurrencyDto) (*models.Currency, error)
}

type ExternalCurrencyApi interface {
	GetCurrenciesByDate(slug []string, date time.Time) (*currencyapi_com.CurrenciesComResponse, error)
	GetCurrenciesFromTo(base string, currencies []string, date time.Time) (*currencyapi_com.CurrenciesComResponse, error)
}

type CurrencyRepository struct {
	storage CurrencyStorage
	api     ExternalCurrencyApi
}

func NewCurrencyRepository(storage CurrencyStorage, api ExternalCurrencyApi) *CurrencyRepository {
	return &CurrencyRepository{storage: storage, api: api}
}

func (r *CurrencyRepository) GetBySlug(ctx context.Context, slug []string, date time.Time) ([]models.Currency, error) {
	log.Printf("CurrencyRepository.GetBySlug slug: %s, date: %s", slug, date)

	curs, err := r.storage.GetBySlug(ctx, slug, date)
	if err != nil {
		log.Printf("CurrencyRepository.GetBySlug, err: %v", err)
		return nil, fmt.Errorf("get currency due err: %w", err)
	}

	if len(curs) == 0 {

		res, err := r.api.GetCurrenciesByDate(slug, date)
		if err != nil {
			return nil, fmt.Errorf("get currency due err: %w", err)
		}

		currencies := make([]models.Currency, 0, len(slug))
		for _, slug := range slug {
			newCurrency := &models.CreateCurrencyDto{
				Slug:  slug,
				Value: res.Data[slug].Value,
				Date:  date,
				Base:  "USD",
			}

			currency, err := r.storage.Create(ctx, newCurrency)
			if err != nil {
				return nil, fmt.Errorf("get currency due err: %w", err)
			}

			currencies = append(currencies, *currency)
		}

		return currencies, nil
	}

	return curs, nil
}

func (r *CurrencyRepository) GetBySlugAndBase(ctx context.Context, slugs []string, base string, date time.Time) ([]models.Currency, error) {
	log.Printf("CurrencyRepository.GetBySlugAndBase slug: %s, base: %s, date: %s", slugs, base, date)

	currencies, err := r.storage.GetBySlugAndBase(ctx, slugs, base, date)
	if err != nil {
		log.Printf("CurrencyRepository.GetBySlugAndBase, err: %v", err)
		return nil, fmt.Errorf("get currency due err: %w", err)
	}

	if len(currencies) == 0 {
		res, err := r.api.GetCurrenciesFromTo(base, slugs, date)
		if err != nil {
			return nil, fmt.Errorf("get currency due err: %w", err)
		}

		currencies := make([]models.Currency, 0, len(slugs))

		for _, slug := range slugs {
			newCurrency := &models.CreateCurrencyDto{
				Slug:  slug,
				Value: res.Data[slug].Value,
				Date:  date,
				Base:  base,
			}

			currency, err := r.storage.Create(ctx, newCurrency)
			if err != nil {
				return nil, fmt.Errorf("get currency due err: %w", err)
			}

			currencies = append(currencies, *currency)
		}

		return currencies, nil

	}

	return currencies, nil
}
