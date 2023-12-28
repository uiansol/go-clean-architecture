package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uiansol/go-clean-architecture/internal/domain/usecases"
)

type UserCreateHandler struct {
	userCreateUseCase usecases.UserCreateUseCase
}

func NewUserCreateHandler(userCreateUseCase usecases.UserCreateUseCase) *UserCreateHandler {
	return &UserCreateHandler{
		userCreateUseCase: userCreateUseCase,
	}
}

func (h *UserCreateHandler) Handle(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
