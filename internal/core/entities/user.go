package entities

import (
	"errors"
	"net/mail"

	"github.com/google/uuid"
)

type User struct {
	ID    string
	Name  string
	Email string
}

func NewUser(name string, address string) (User, error) {

	email, err := validMailAddress(address)
	if err != nil {
		return User{}, errors.New("invalid email address")
	}

	uuid := uuid.NewString()

	user := User{
		ID:    uuid,
		Name:  name,
		Email: email,
	}

	return user, nil
}

func validMailAddress(address string) (string, error) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", err
	}

	return addr.Address, nil
}
