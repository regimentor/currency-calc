package http

import (
	"fmt"
	"github.com/regimentor/currency-calc/internal"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	repository UserRepository
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	newUser, err := h.repository.Create(internal.CreateUserDto{})
	if err != nil {
		return fmt.Errorf("crete user due err: %v", err)
	}

	return c.JSON(http.StatusOK, newUser)

}
