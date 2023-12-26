package repositories

import "github.com/uiansol/go-clean-architecture/internal/core/entities"

type IUserRepository interface {
	Save(user entities.User) (string, error)
}
