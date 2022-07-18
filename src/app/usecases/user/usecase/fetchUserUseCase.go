package usecase

import (
	"my-module/app/domain/user/repository"
	"my-module/app/errors"
	"my-module/app/usecases/user/dto"
)

type FetchUserUseCase struct {
	ur repository.IUserRepository
}

func NewFetchUserUseCase(ur repository.IUserRepository) *FetchUserUseCase {
	return &FetchUserUseCase{ur}
}

func (uc *FetchUserUseCase) Execute(userId int) (*dto.UserDto, *errors.AppError) {

	ue, err := uc.ur.FindById(userId)
	if err.HasErrors() {
		return nil, err
	}

	userDto := dto.NewUserDto(
		ue.GetId(),
		ue.GetUserSubId().GetValue(),
		ue.GetMailAddress().GetValue(),
	)

	return userDto, nil
}
