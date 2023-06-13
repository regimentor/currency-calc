package http

import (
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/regimentor/currency-calc/internal"
)

type ApiLogsRepository interface {
	Create(apiLogs *internal.CreateApiLogsDto) error
}

type UserRepository interface {
	Create(u internal.CreateUserDto) (*internal.User, error)
	GetByApiKey(apiKey internal.ApiKey) (*internal.User, error)
	GetById(id internal.UserId) (*internal.User, error)
}

type CurrenciesRepository interface {
	GetBySlug(currencies []string, date time.Time) ([]internal.Currency, error)
	GetBySlugAndBase(slugs []string, base string, date time.Time) ([]internal.Currency, error)
}

type Server struct {
	userRepository       UserRepository
	currenciesRepository CurrenciesRepository
	apiLogsRepository    ApiLogsRepository
}

func NewServer(userRepository UserRepository, currenciesRepository CurrenciesRepository, apiLogsRepository ApiLogsRepository) *Server {
	return &Server{
		userRepository:       userRepository,
		currenciesRepository: currenciesRepository,
		apiLogsRepository:    apiLogsRepository,
	}
}

func AuthenticationMiddleware(userRepository UserRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Printf("AuthenticationMiddleware, url: %s", c.Request().URL)
			apiKey := c.Request().Header.Get("X-API-KEY")
			if apiKey == "" {
				return echo.ErrUnauthorized
			}

			user, err := userRepository.GetByApiKey(internal.ApiKey(apiKey))
			if err != nil {
				return echo.ErrUnauthorized
			}

			c.Set("userId", user.ID)

			return next(c)
		}
	}
}

func (s *Server) Listen() error {
	server := echo.New()
	authorised := server.Group("/api", AuthenticationMiddleware(s.userRepository))

	userHandler := UserHandler{repository: s.userRepository}
	server.POST("/user", userHandler.CreateUser)

	currencyHandler := CurrencyHandler{
		repository:        s.currenciesRepository,
		apiLogsRepository: s.apiLogsRepository,
	}
	authorised.GET("/currencies", currencyHandler.GetCurrencies)
	authorised.GET("/currencies/from-to", currencyHandler.GetCurrenciesFromTo)

	err := server.Start(":8080")
	if err != nil {
		return fmt.Errorf("starting server due err: %v", err)
	}

	return nil
}
