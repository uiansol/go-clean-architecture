package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/go-clean-architecture/internal/domain/usecases"
)

func TestHandle(t *testing.T) {
	t.Run("should process data and return error nil", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/user?name=test&email=test@test.com", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := usecases.NewIUserCreateUseCase(t)
		useCaseMock.On("Execute", mock.Anything).Return(usecases.UserCreateOutput{UserID: "Test-ID"}, nil)

		h := NewUserCreateHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, "{\"user_id\":\"Test-ID\"}", rec.Body.String())
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
