package usecase

import (
	"my-module/app/domain/player/repository"
	vo "my-module/app/domain/player/valueObject"
	"my-module/app/errors"
	"my-module/app/usecases/player/dto"
)

type FetchAllPlayersByContestIdUseCase struct {
	r repository.IPlayerRepository
}

func NewFetchAllPlayersByContestIdUseCase(r repository.IPlayerRepository) *FetchAllPlayersByContestIdUseCase {
	return &FetchAllPlayersByContestIdUseCase{r}
}

func (uc *FetchAllPlayersByContestIdUseCase) Execute(contestId int) ([]dto.PlayerDto, *errors.AppError) {

	// ソートタイプを降順へ指定
	sortType, err := vo.NewSortType(vo.SortEnum("Desc"))
	if err.HasErrors() {
		return nil, err
	}

	playerEntities, err := uc.r.GetAllByContestId(contestId, sortType)
	if err.HasErrors() {
		return nil, err
	}

	var playerDtos []dto.PlayerDto
	var playerDto dto.PlayerDto
	for _, playerEntity := range playerEntities {
		playerDto = *dto.NewPlayerDto(
			playerEntity.GetId(),
			playerEntity.GetUserEntity().GetUserSubId().GetValue(),
			playerEntity.GetContestEntity().GetId(),
			playerEntity.GetName().GetValue(),
			playerEntity.GetImagePath().GetValue(),
			playerEntity.GetDescription().GetValue(),
			playerEntity.GetScore().GetValue(),
			playerEntity.GetClicked().GetValue(),
		)
		playerDtos = append(playerDtos, playerDto)
	}

	return playerDtos, nil
}
