package usecase

import (
	"my-module/app/domain/contest/repository"
	"my-module/app/errors"
	"my-module/app/usecases/contest/dto"
)

type FetchAllContestsUseCase struct {
	r repository.IContestRepository
}

func NewFetchAllContestsUseCase(r repository.IContestRepository) *FetchAllContestsUseCase {
	return &FetchAllContestsUseCase{r}
}

func (uc *FetchAllContestsUseCase) Execute() ([]dto.ContestDto, *errors.AppError) {
	contestEntities, err := uc.r.GetAll()
	if err.HasErrors() {
		return nil, err
	}

	var contestDtos []dto.ContestDto
	var contestDto dto.ContestDto
	for _, contestEntity := range contestEntities {
		contestDto = *dto.NewContestDto(
			contestEntity.GetId(),
			contestEntity.GetUserEntity().GetUserSubId().GetValue(),
			contestEntity.GetName().GetValue(),
			contestEntity.GetDescription().GetValue(),
			contestEntity.GetImagePath().GetValue(),
		)
		contestDtos = append(contestDtos, contestDto)
	}

	return contestDtos, nil
}
