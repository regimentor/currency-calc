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

func (h *CurrencyHandler) GetCurrencies(c echo.Context) error {
	log.Println("_____________________get currencies______________________")

	if err := h.apiLogsRepository.Create(&internal.CreateApiLogsDto{
		UserId:      c.Get("userId").(internal.UserId),
		RequestType: internal.GET_ALL,
		RequestTime: time.Now(),
	}); err != nil {
		log.Printf("CurrencyHandler.GetCurrencies create api logs due err: %v", err)
	}

	date := c.QueryParam("date")
	currencies := strings.Split(c.QueryParam("currencies"), ",")

	parsedTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Printf("CurrencyHandler.GetCurrencies parse time due err: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	byDate, err := h.repository.GetBySlug(currencies, parsedTime)
	if err != nil {
		log.Printf("CurrencyHandler.GetCurrencies get currencies due err: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, byDate)
}

func (h *CurrencyHandler) GetCurrenciesFromTo(c echo.Context) error {
	log.Println("_____________________get currencies from to______________________")

	if err := h.apiLogsRepository.Create(&internal.CreateApiLogsDto{
		UserId:      c.Get("userId").(internal.UserId),
		RequestType: internal.GET_BY_BASE,
		RequestTime: time.Now(),
	}); err != nil {
		log.Printf("CurrencyHandler.GetCurrenciesFromTo create api logs due err: %v", err)
	}

	date := c.QueryParam("date")
	currencies := strings.Split(c.QueryParam("currencies"), ",")
	base := c.QueryParam("base")

	parsedTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Printf("CurrencyHandler.GetCurrenciesFromTo parse time due err: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	byDate, err := h.repository.GetBySlugAndBase(currencies, base, parsedTime)
	if err != nil {
		log.Printf("CurrencyHandler.GetCurrenciesFromTo get currencies due err: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, byDate)
}
