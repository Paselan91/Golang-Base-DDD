package usecase

import (
	// "my-module/app/domain/user/entity"
	"my-module/app/domain/user/factory"
	"my-module/app/domain/user/repository"
	"my-module/app/errors"
	"my-module/app/http/inputs/userInputs"
	"my-module/app/usecases/user/dto"
)

type CreateUserUseCase struct {
	ur repository.IUserRepository
}

func NewCreateUserUseCase(ur repository.IUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{ur}
}

func (uc *CreateUserUseCase) Execute(input *userInputs.CreateUserInput) (*dto.UserDto, *errors.AppError) {

	uf := factory.NewUserFactory()
	userEntity, err := uf.Create(
		input.GetUserSubId(),
		input.GetMailAddress(),
	)
	if err.HasErrors() {
		return nil, err
	}

	ue, err := uc.ur.Create(userEntity)
	if err.HasErrors() {
		return nil, err
	}

	u := dto.UserDto{
		Id:          ue.GetId(),
		UserSubId:   ue.GetUserSubId().GetValue(),
		MailAddress: ue.GetMailAddress().GetValue(),
	}

	return &u, nil
}
