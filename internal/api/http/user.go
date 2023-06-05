package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/regimentor/currency-calc/internal"
)

type UserHandler struct {
	repository UserRepository
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	newUser, err := h.repository.Create(&internal.User{})
	if err != nil {
		return fmt.Errorf("crete user due err: %v", err)
	}

	return c.JSON(http.StatusOK, newUser)

}
