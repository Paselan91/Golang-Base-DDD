package repository

import (
	"fmt"
	"gorm.io/gorm"
	"my-module/app/domain/player/entity"
	"my-module/app/domain/player/repository"
	vo "my-module/app/domain/player/valueObject"
	"my-module/app/errors"
	"my-module/app/infrastructure/models"
	// "log"
)

type playerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) repository.IPlayerRepository {
	return &playerRepository{db}
}

func (r *playerRepository) GetAllByContestId(
	contestId int,
	sortType *vo.SortType,
) ([]*entity.PlayerEntity, *errors.AppError) {

	players := []models.Player{}

	orderKey, errS := getOrderKey(sortType)
	if errS != nil {
		errS := errors.NewAppError("対象コンテストの参加者が見つかりませんでした。")
		return nil, &errS
	}

	err := r.db.Preload("User").Preload("Contest").Order(orderKey).Find(&players).Error
	if err != nil {
		err := errors.NewAppError("対象コンテストの参加者が見つかりませんでした。")
		return nil, &err
	}
	var playerEntities []*entity.PlayerEntity

	for _, player := range players {
		playerEntity, errCe := player.ToEntity()
		if errCe.HasErrors() {
			return nil, errCe
		}

		playerEntities = append(playerEntities, playerEntity)
	}

	return playerEntities, nil
}

func getOrderKey(s *vo.SortType) (string, *errors.AppError) {
	sortType := s.GetValue()
	switch sortType {
	case vo.Desc:
		return "created_at desc", nil
	case vo.ScoreDesc:
		return "score desc", nil
	default:
		msg := fmt.Sprintf("ソートタイプkeyへ不正な値が指定されました。sortType: %s\n", sortType)
		err := errors.NewAppError(msg)
		return "err", &err
	}
}
