package http

import (
	"github.com/labstack/echo/v4"
	"github.com/regimentor/currency-calc/internal"
	"log"
	"net/http"
	"strings"
	"time"
)

type CurrencyHandler struct {
	repository        CurrenciesRepository
	apiLogsRepository ApiLogsRepository
}

func inArray(target string, array []string) bool {
	for _, value := range array {
		if value == target {
			return true
		}
	}
	return false
}

func validateSlugs(slugs []string) bool {
	validSlugs := []string{"USD", "EUR", "RUB", "JPY"}
	for _, slug := range slugs {
		if !inArray(slug, validSlugs) {
			return false
		}
	}

	return true
}

func (h *CurrencyHandler) GetCurrencies(c echo.Context) error {
	log.Println("_____________________get currencies______________________")
	ctx := c.Request().Context()

	if err := h.apiLogsRepository.Create(ctx, &internal.CreateApiLogsDto{
		UserId:      c.Get("userId").(internal.UserId),
		RequestType: internal.GetAll,
		RequestTime: time.Now(),
	}); err != nil {
		log.Printf("CurrencyHandler.GetCurrencies create api logs due err: %v", err)
	}

	currencies := strings.Split(c.QueryParam("currencies"), ",")
	if !validateSlugs(currencies) {
		log.Printf("CurrencyHandler.GetCurrencies invalid currencies: %v", currencies)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid currencies")
	}

	date := c.QueryParam("date")
	parsedTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Printf("CurrencyHandler.GetCurrencies parse time due err: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid date")
	}

	byDate, err := h.repository.GetBySlug(ctx, currencies, parsedTime)
	if err != nil {
		log.Printf("CurrencyHandler.GetCurrencies get currencies due err: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, byDate)
}

func (h *CurrencyHandler) GetCurrenciesFromTo(c echo.Context) error {
	log.Println("_____________________get currencies from to______________________")
	ctx := c.Request().Context()

	if err := h.apiLogsRepository.Create(ctx, &internal.CreateApiLogsDto{
		UserId:      c.Get("userId").(internal.UserId),
		RequestType: internal.GetByBase,
		RequestTime: time.Now(),
	}); err != nil {
		log.Printf("CurrencyHandler.GetCurrenciesFromTo create api logs due err: %v", err)
	}

	currencies := strings.Split(c.QueryParam("currencies"), ",")
	if !validateSlugs(currencies) {
		log.Printf("CurrencyHandler.GetCurrencies invalid currencies: %v", currencies)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid currencies")
	}

	date := c.QueryParam("date")
	base := c.QueryParam("base")

	parsedTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Printf("CurrencyHandler.GetCurrenciesFromTo parse time due err: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid date")
	}

	byDate, err := h.repository.GetBySlugAndBase(ctx, currencies, base, parsedTime)
	if err != nil {
		log.Printf("CurrencyHandler.GetCurrenciesFromTo get currencies due err: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, byDate)
}
