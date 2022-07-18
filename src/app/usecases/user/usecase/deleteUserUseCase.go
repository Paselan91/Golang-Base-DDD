package usecase

import (
	"my-module/app/domain/user/repository"
	"my-module/app/errors"
)

type DeleteUserUseCase struct {
	ur repository.IUserRepository
}

func NewDeleteUserUseCase(ur repository.IUserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{ur}
}

func (uc *DeleteUserUseCase) Execute(userId int) (bool, *errors.AppError) {

	ok, err := uc.ur.Delete(userId)
	if err.HasErrors() {
		return false, err
	}

	return ok, nil
}
