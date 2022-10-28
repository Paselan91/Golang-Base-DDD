package repository

import (
	"my-module/app/domain/user/entity"
	"my-module/app/errors"
)

type IUserRepository interface {
	FindById(id int) (*entity.UserEntity, *errors.AppError)
	Create(userEntity *entity.UserEntity) (*entity.UserEntity, *errors.AppError)
	Delete(id int) (bool, *errors.AppError)
}
