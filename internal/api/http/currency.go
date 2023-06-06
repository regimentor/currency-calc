package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type CurrencyHandler struct {
}

func (h *CurrencyHandler) GetCurrencies(c echo.Context) error {
	return c.String(http.StatusOK, "/currencies")
}

func (h *CurrencyHandler) GetCurrenciesFromTo(c echo.Context) error {
	return c.String(http.StatusOK, "/currencies/from-to")
}
