package repository

import (
	"gorm.io/gorm"
	"my-module/app/domain/contest/entity"
	"my-module/app/domain/contest/repository"
	"my-module/app/errors"
	"my-module/app/infrastructure/models"
)

type contestRepository struct {
	db *gorm.DB
}

func NewContestRepository(db *gorm.DB) repository.IContestRepository {
	return &contestRepository{db}
}

func (r *contestRepository) GetAll() ([]*entity.ContestEntity, *errors.AppError) {

	contests := []models.Contest{}

	err := r.db.Preload("User").Find(&contests).Error
	if err != nil {
		err := errors.NewAppError("コンテストが見つかりませんでした。")
		return nil, &err
	}
	var contestEntities []*entity.ContestEntity

	for _, contest := range contests {
		contestEntity, errCe := contest.ToEntity()
		if errCe.HasErrors() {
			return nil, errCe
		}
		contestEntities = append(contestEntities, contestEntity)
	}

	return contestEntities, nil
}
