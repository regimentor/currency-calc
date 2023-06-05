package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/regimentor/currency-calc/internal"
	"net/http"
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

func (s *Server) Listen() error {
	server := echo.New()

	userHandler := UserHandler{repository: s.userRepository}
	server.POST("/user", userHandler.CreateUser)

	authorised := server.Group("/api", Authentication)
	authorised.GET("/currencies", GetCurrencies)
	authorised.GET("/currencies/from-to", GetCurrenciesFromTo)

	err := server.Start(":8080")
	if err != nil {
		return fmt.Errorf("starting server due err: %v", err)
	}

	return nil
}
