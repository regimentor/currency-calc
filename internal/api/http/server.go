package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/regimentor/currency-calc/internal"
)

type UserRepository interface {
	Create(u internal.CreateUserDto) (*internal.User, error)
}

type Server struct {
	userRepository UserRepository
}

func NewServer(userRepository UserRepository) *Server {
	return &Server{userRepository: userRepository}
}

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		return next(c)
	}
}

func (s *Server) Listen() error {
	server := echo.New()
	authorised := server.Group("/api", Authentication)

	userHandler := UserHandler{repository: s.userRepository}
	server.POST("/user", userHandler.CreateUser)

	currencyHandler := CurrencyHandler{}
	authorised.GET("/currencies", currencyHandler.GetCurrencies)
	authorised.GET("/currencies/from-to", currencyHandler.GetCurrenciesFromTo)

	err := server.Start(":8080")
	if err != nil {
		return fmt.Errorf("starting server due err: %v", err)
	}

	return nil
}
