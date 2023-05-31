package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func CreateToken(c echo.Context) error {

	return c.String(http.StatusOK, "/create-token")
}

func GetCurrencies(c echo.Context) error {

	return c.String(http.StatusOK, "/currencies")
}

func GetCurrenciesFromTo(c echo.Context) error {

	return c.String(http.StatusOK, "/currencies/from-to")
}

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		return next(c)
	}
}

func (s Server) Listen() error {
	server := echo.New()

	server.POST("/create-token", CreateToken)

	authorised := server.Group("/api", Authentication)
	authorised.GET("/currencies", GetCurrencies)
	authorised.GET("/currencies/from-to", GetCurrenciesFromTo)

	err := server.Start(":8080")
	if err != nil {
		return fmt.Errorf("starting server due err: %v", err)
	}

	return nil
}
