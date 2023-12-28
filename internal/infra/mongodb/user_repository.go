package mongodb

import "github.com/uiansol/go-clean-architecture/internal/domain/entities"

type UserRepositoryMongo struct {
	// Some variables here to manipulate mongodb
}

func (r *UserRepositoryMongo) Save(user entities.User) (string, error) {
	// Implements logic for save on mongodb
	return "", nil
}
