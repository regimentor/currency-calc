package internal

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/regimentor/currency-calc/internal/currencyapi.com"
	"github.com/regimentor/currency-calc/internal/models"
	"github.com/regimentor/currency-calc/mocks"
	"testing"
	"time"
)

func TestCurrencyRepository_GetBySlug(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	currencyStorage := mock_internal.NewMockCurrencyStorage(ctrl)
	currencyApi := mock_internal.NewMockExternalCurrencyApi(ctrl)
	currencyRepository := NewCurrencyRepository(currencyStorage, currencyApi)

	slugs := []string{"USD", "EUR"}
	date := time.Now()

	exceptedCurrency := []models.Currency{
		models.Currency{},
		models.Currency{},
	}

	data := map[string]currencyapi_com.ResponseCurrenciesData{
		"USD": {
			Code:  "USD",
			Value: 1.0,
		},
		"EUR": {
			Code:  "EUR",
			Value: 1.0,
		},
	}

	exceptedCurrencyApi := &currencyapi_com.CurrenciesComResponse{
		Data: data,
	}

	exceptedCurrencyEmpty := make([]models.Currency, 0)

	t.Run("Get currency by slug", func(t *testing.T) {

		currencyStorage.
			EXPECT().
			GetBySlug(context.Background(), slugs, date).
			Return(exceptedCurrency, nil).
			Times(1)

		_, err := currencyRepository.GetBySlug(
			context.Background(), slugs, date)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("Get currency by slug currency repository due err", func(t *testing.T) {
		currencyStorage.
			EXPECT().
			GetBySlug(context.Background(), slugs, date).
			Return(nil, errors.New("get currencies due err")).
			Times(1)

		_, err := currencyRepository.GetBySlug(
			context.Background(), slugs, date)

		if err == nil {
			t.Errorf("expected error: %v", err)
		}
	})

	t.Run("Get currency by slug currency repository due empty slice", func(t *testing.T) {
		currencyStorage.
			EXPECT().
			GetBySlug(context.Background(), slugs, date).
			Return(exceptedCurrencyEmpty, nil).
			Times(1)

		currencyApi.
			EXPECT().
			GetCurrenciesByDate(slugs, date).
			Return(exceptedCurrencyApi, nil).
			Times(1)

		currencyStorage.
			EXPECT().
			Create(context.Background(), &models.CreateCurrencyDto{
				Slug:  "USD",
				Value: 1.0,
				Date:  date,
				Base:  "USD",
			}).
			Return(&models.Currency{Slug: "USD",
				Value: 1.0,
				Date:  date,
				Base:  "USD"}, nil).
			Times(1)

		currencyStorage.
			EXPECT().
			Create(context.Background(), &models.CreateCurrencyDto{
				Slug:  "EUR",
				Value: 1.0,
				Date:  date,
				Base:  "USD",
			}).
			Return(&models.Currency{Slug: "EUR",
				Value: 1.0,
				Date:  date,
				Base:  "USD"}, nil).
			Times(1)

		_, err := currencyRepository.GetBySlug(
			context.Background(), slugs, date)

		if err != nil {
			t.Errorf("expected error: %v", err)
		}
	})

	t.Run("currencyApi.GetCurrenciesByDate due err", func(t *testing.T) {
		currencyStorage.
			EXPECT().
			GetBySlug(context.Background(), slugs, date).
			Return(exceptedCurrencyEmpty, nil).
			Times(1)

		currencyApi.
			EXPECT().
			GetCurrenciesByDate(slugs, date).
			Return(nil, errors.New("currencyApi.GetCurrenciesByDate due err")).
			Times(1)

		_, err := currencyRepository.GetBySlug(
			context.Background(), slugs, date)

		if err == nil {
			t.Errorf("expected error: %v", err)
		}
	})

	t.Run("currencyStorage.Create due err", func(t *testing.T) {
		currencyStorage.
			EXPECT().
			GetBySlug(context.Background(), slugs, date).
			Return(exceptedCurrencyEmpty, nil).
			Times(1)

		currencyApi.
			EXPECT().
			GetCurrenciesByDate(slugs, date).
			Return(exceptedCurrencyApi, nil).
			Times(1)

		currencyStorage.
			EXPECT().
			Create(context.Background(), &models.CreateCurrencyDto{
				Slug:  "USD",
				Value: 1.0,
				Date:  date,
				Base:  "USD",
			}).
			Return(nil, errors.New("currencyStorage.Create due err")).
			Times(1)

		_, err := currencyRepository.GetBySlug(
			context.Background(), slugs, date)

		if err == nil {
			t.Errorf("expected error: %v", err)
		}
	})
}

func TestCurrencyRepository_GetBySlugAndDate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	currencyStorage := mock_internal.NewMockCurrencyStorage(ctrl)
	currencyApi := mock_internal.NewMockExternalCurrencyApi(ctrl)
	currencyRepository := NewCurrencyRepository(currencyStorage, currencyApi)

	slugs := []string{"USD", "EUR"}
	date := time.Now()
	base := "USD"

	exceptedCurrency := []models.Currency{
		models.Currency{},
		models.Currency{},
	}

	data := map[string]currencyapi_com.ResponseCurrenciesData{
		"USD": {
			Code:  "USD",
			Value: 1.0,
		},
		"EUR": {
			Code:  "EUR",
			Value: 1.0,
		},
	}

	exceptedCurrencyApi := &currencyapi_com.CurrenciesComResponse{
		Data: data,
	}

	exceptedCurrencyEmpty := make([]models.Currency, 0)

	t.Run("Get currency by slug and base", func(t *testing.T) {

		currencyStorage.
			EXPECT().
			GetBySlugAndBase(context.Background(), slugs, base, date).
			Return(exceptedCurrency, nil).
			Times(1)

		_, err := currencyRepository.GetBySlugAndBase(
			context.Background(), slugs, base, date)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("Get currency by slug currency repository due err", func(t *testing.T) {
		currencyStorage.
			EXPECT().
			GetBySlugAndBase(context.Background(), slugs, base, date).
			Return(nil, errors.New("get currencies due err")).
			Times(1)

		_, err := currencyRepository.GetBySlugAndBase(
			context.Background(), slugs, base, date)

		if err == nil {
			t.Errorf("expected error: %v", err)
		}
	})

	t.Run("Get currency by slug currency repository due empty slice", func(t *testing.T) {
		currencyStorage.
			EXPECT().
			GetBySlugAndBase(context.Background(), slugs, base, date).
			Return(exceptedCurrencyEmpty, nil).
			Times(1)

		currencyApi.
			EXPECT().
			GetCurrenciesFromTo(base, slugs, date).
			Return(exceptedCurrencyApi, nil).
			Times(1)

		currencyStorage.
			EXPECT().
			Create(context.Background(), &models.CreateCurrencyDto{
				Slug:  "USD",
				Value: 1.0,
				Date:  date,
				Base:  "USD",
			}).
			Return(&models.Currency{Slug: "USD",
				Value: 1.0,
				Date:  date,
				Base:  "USD"}, nil).
			Times(1)

		currencyStorage.
			EXPECT().
			Create(context.Background(), &models.CreateCurrencyDto{
				Slug:  "EUR",
				Value: 1.0,
				Date:  date,
				Base:  "USD",
			}).
			Return(&models.Currency{Slug: "EUR",
				Value: 1.0,
				Date:  date,
				Base:  "USD"}, nil).
			Times(1)

		_, err := currencyRepository.GetBySlugAndBase(
			context.Background(), slugs, base, date)

		if err != nil {
			t.Errorf("expected error: %v", err)
		}
	})

	t.Run("currencyApi.GetCurrenciesByDate due err", func(t *testing.T) {
		currencyStorage.
			EXPECT().
			GetBySlugAndBase(context.Background(), slugs, base, date).
			Return(exceptedCurrencyEmpty, nil).
			Times(1)

		currencyApi.
			EXPECT().
			GetCurrenciesFromTo(base, slugs, date).
			Return(nil, errors.New("currencyApi.GetCurrenciesByDate due err")).
			Times(1)

		_, err := currencyRepository.GetBySlugAndBase(
			context.Background(), slugs, base, date)

		if err == nil {
			t.Errorf("expected error: %v", err)
		}
	})

	t.Run("currencyStorage.Create due err", func(t *testing.T) {
		currencyStorage.
			EXPECT().
			GetBySlugAndBase(context.Background(), slugs, base, date).
			Return(exceptedCurrencyEmpty, nil).
			Times(1)

		currencyApi.
			EXPECT().
			GetCurrenciesFromTo(base, slugs, date).
			Return(exceptedCurrencyApi, nil).
			Times(1)

		currencyStorage.
			EXPECT().
			Create(context.Background(), &models.CreateCurrencyDto{
				Slug:  "USD",
				Value: 1.0,
				Date:  date,
				Base:  "USD",
			}).
			Return(nil, errors.New("currencyStorage.Create due err")).
			Times(1)

		_, err := currencyRepository.GetBySlugAndBase(
			context.Background(), slugs, base, date)

		if err == nil {
			t.Errorf("expected error: %v", err)
		}
	})
}
