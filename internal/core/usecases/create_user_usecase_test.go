package usecases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/go-clean-architecture/mocks"
)

func TestNewCreateUserUseCase(t *testing.T) {
	userRepositoryMock := mocks.NewIUserRepository(t)

	t.Run("should return a create user use case", func(t *testing.T) {
		createUserUseCase := NewCreateUserUseCase(userRepositoryMock)

		assert.NotNil(t, createUserUseCase)
	})
}

func TestExecute(t *testing.T) {
	input := CreateUserInput{
		Name:  "test user",
		Email: "testuser@testuser.com",
	}

	t.Run("should create the user and return error nil", func(t *testing.T) {
		userRepositoryMock := mocks.NewIUserRepository(t)
		createUserUseCase := NewCreateUserUseCase(userRepositoryMock)

		userRepositoryMock.On("Save", mock.Anything).Return("123", nil)

		output, err := createUserUseCase.Execute(input)

		assert.Nil(t, err)
		assert.Equal(t, "123", output.UserID)
	})

	t.Run("should return error with email is invalid", func(t *testing.T) {
		userRepositoryMock := mocks.NewIUserRepository(t)
		createUserUseCase := NewCreateUserUseCase(userRepositoryMock)

		inputEmailInvalid := CreateUserInput{
			Name:  "test user",
			Email: "testusertestuser.com",
		}

		output, err := createUserUseCase.Execute(inputEmailInvalid)

		assert.Equal(t, "invalid email address", err.Error())
		assert.Equal(t, CreateUserOutput{}, output)
	})

	t.Run("should return error when couldn't save", func(t *testing.T) {
		userRepositoryMock := mocks.NewIUserRepository(t)
		createUserUseCase := NewCreateUserUseCase(userRepositoryMock)

		userRepositoryMock.On("Save", mock.Anything).Return("", errors.New("test error"))

		output, err := createUserUseCase.Execute(input)
		assert.Equal(t, "test error", err.Error())
		assert.Equal(t, CreateUserOutput{}, output)
	})
}
