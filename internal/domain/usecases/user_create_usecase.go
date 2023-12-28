package usecases

import (
	"github.com/uiansol/go-clean-architecture/internal/domain/adapters"
	"github.com/uiansol/go-clean-architecture/internal/domain/entities"
)

type UserCreateUseCase struct {
	userRepository adapters.IUserRepository
}

type UserCreateInput struct {
	Name  string
	Email string
}

type UserCreateOutput struct {
	UserID string
}

func NewUserCreateUseCase(userRepository adapters.IUserRepository) UserCreateUseCase {
	return UserCreateUseCase{
		userRepository: userRepository,
	}
}

func (u *UserCreateUseCase) Execute(userCreateInput UserCreateInput) (UserCreateOutput, error) {
	user, err := entities.NewUser(userCreateInput.Name, userCreateInput.Email)
	if err != nil {
		return UserCreateOutput{}, err
	}

	id, err := u.userRepository.Save(user)
	if err != nil {
		return UserCreateOutput{}, err
	}

	output := UserCreateOutput{
		UserID: id,
	}

	return output, nil
}
