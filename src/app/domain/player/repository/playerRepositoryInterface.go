package repository

import (
	"my-module/app/domain/player/entity"
	vo "my-module/app/domain/player/valueObject"
	"my-module/app/errors"
)

type IPlayerRepository interface {
	GetAllByContestId(contestId int, sortType *vo.SortType) ([]*entity.PlayerEntity, *errors.AppError)
}
