package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/regimentor/currency-calc/internal"
	"log"
)

type UserRepository interface {
	Create(u internal.CreateUserDto) (*internal.User, error)
	GetByApiKey(apiKey internal.ApiKey) (*internal.User, error)
	GetById(id internal.UserId) (*internal.User, error)
}

type Server struct {
	userRepository UserRepository
}

func NewServer(userRepository UserRepository) *Server {
	return &Server{userRepository: userRepository}
}

func AuthenticationMiddleware(userRepository UserRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Printf("AuthenticationMiddleware, url: %s", c.Request().URL)
			apiKey := c.Request().Header.Get("X-API-KEY")
			if apiKey == "" {
				return echo.ErrUnauthorized
			}

			_, err := userRepository.GetByApiKey(internal.ApiKey(apiKey))
			if err != nil {
				return echo.ErrUnauthorized
			}

			return next(c)
		}
	}
}

func (s *Server) Listen() error {
	server := echo.New()
	authorised := server.Group("/api", AuthenticationMiddleware(s.userRepository))

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
