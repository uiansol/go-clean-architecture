package api

import (
	"github.com/uiansol/go-clean-architecture/internal/domain/usecases"
	"github.com/uiansol/go-clean-architecture/internal/infra/api/handlers"
	"github.com/uiansol/go-clean-architecture/internal/infra/mongodb"
)

func configHandlers() *AppHandlers {
	userRepository := mongodb.NewUserRepositoryMongo()
	userCreateUseCase := usecases.NewUserCreateUseCase(userRepository)
	userCreateHandler := handlers.NewUserCreateHandler(&userCreateUseCase)

	return &AppHandlers{
		userCreateHandler: userCreateHandler,
		// Add new handlers here
	}
}
