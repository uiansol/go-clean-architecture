package mappers

import (
	"github.com/uiansol/go-clean-architecture/internal/domain/usecases"
	"github.com/uiansol/go-clean-architecture/internal/infra/api/dto"
)

func UserCreateRequestToUserCreateInput(u dto.UserCreateRequest) usecases.UserCreateInput {
	return usecases.UserCreateInput{
		Name:  u.Name,
		Email: u.Email,
	}
}

func UserCreateOutputToUserCreateResponse(u usecases.UserCreateOutput) dto.UserCreateResponse {
	return dto.UserCreateResponse{
		UserID: u.UserID,
	}
}
