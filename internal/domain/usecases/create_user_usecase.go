package usecases

import (
	"github.com/uiansol/go-clean-architecture/internal/domain/adapters"
	"github.com/uiansol/go-clean-architecture/internal/domain/entities"
)

type CreateUserUseCase struct {
	userRepository adapters.IUserRepository
}

type CreateUserInput struct {
	Name  string
	Email string
}

type CreateUserOutput struct {
	UserID string
}

func NewCreateUserUseCase(userRepository adapters.IUserRepository) CreateUserUseCase {
	return CreateUserUseCase{
		userRepository: userRepository,
	}
}

func (u *CreateUserUseCase) Execute(createUserInput CreateUserInput) (CreateUserOutput, error) {
	user, err := entities.NewUser(createUserInput.Name, createUserInput.Email)
	if err != nil {
		return CreateUserOutput{}, err
	}

	id, err := u.userRepository.Save(user)
	if err != nil {
		return CreateUserOutput{}, err
	}

	output := CreateUserOutput{
		UserID: id,
	}

	return output, nil
}
