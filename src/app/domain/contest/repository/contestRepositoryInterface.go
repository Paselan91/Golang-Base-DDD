package repository

import (
	"my-module/app/domain/contest/entity"
	"my-module/app/errors"
)

type IContestRepository interface {
	GetAll() ([]*entity.ContestEntity, *errors.AppError)
}
