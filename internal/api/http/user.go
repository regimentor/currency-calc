package http

import (
	"fmt"
	"github.com/regimentor/currency-calc/internal"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	repository UserRepository
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	log.Println("create user")
	var apiKey internal.ApiKey
	apiKey.Generate()
	log.Printf("create user %s", apiKey)

	ctx := c.Request().Context()
	newUser, err := h.repository.Create(ctx, internal.CreateUserDto{ApiKey: apiKey})
	if err != nil {
		log.Printf("create user due err: %v", err)
		return fmt.Errorf("crete user due err: %v", err)
	}

	return c.JSON(http.StatusOK, newUser)

}
