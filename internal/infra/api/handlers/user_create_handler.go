package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uiansol/go-clean-architecture/internal/domain/usecases"
	"github.com/uiansol/go-clean-architecture/internal/infra/api/dto"
	"github.com/uiansol/go-clean-architecture/internal/infra/api/mappers"
)

type UserCreateHandler struct {
	userCreateUseCase usecases.IUserCreateUseCase
}

func NewUserCreateHandler(userCreateUseCase usecases.IUserCreateUseCase) *UserCreateHandler {
	return &UserCreateHandler{
		userCreateUseCase: userCreateUseCase,
	}
}

func (h *UserCreateHandler) Handle(c echo.Context) error {
	var user dto.UserCreateRequest

	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	input := mappers.UserCreateRequestToUserCreateInput(user)
	output, err := h.userCreateUseCase.Execute(input)
	if err != nil {
		c.Error(err)
	}

	response := mappers.UserCreateOutputToUserCreateResponse(output)
	responseJson, err := json.Marshal(response)
	if err != nil {
		c.Error(err)
	}

	return c.String(http.StatusOK, string(responseJson))
}
