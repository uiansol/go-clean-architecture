package adapters

import "github.com/uiansol/go-clean-architecture/internal/domain/entities"

type IUserRepository interface {
	Save(user entities.User) (string, error)
}
