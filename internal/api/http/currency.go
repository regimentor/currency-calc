package http

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
	"time"
)

type CurrencyHandler struct {
	repository CurrenciesRepository
}

func (h *CurrencyHandler) GetCurrencies(c echo.Context) error {
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
	return c.String(http.StatusOK, "/currencies/from-to")
}
