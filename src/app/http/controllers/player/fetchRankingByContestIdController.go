package player

import (
	"github.com/labstack/echo"
	"my-module/app/infrastructure/repository"
	"my-module/app/usecases/player/usecase"
	"my-module/config"
	"net/http"
	"strconv"
)

type FetchRankingByContestIdController struct{}

func NewFetchRankingByContestIdController() *FetchRankingByContestIdController {
	return &FetchRankingByContestIdController{}
}

func (controller *FetchRankingByContestIdController) Execute(c echo.Context) error {
	contestId, err := strconv.Atoi(c.Param("contest_id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	db, err := config.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	r := repository.NewPlayerRepository(db)
	fetchRankingByContestIdUseCase := usecase.NewFetchRankingByContestIdUseCase(r)

	playerDtoList, errUc := fetchRankingByContestIdUseCase.Execute(contestId)
	if errUc.HasErrors() {
		return c.JSON(http.StatusInternalServerError, errUc)
	}

	return c.JSON(http.StatusOK, playerDtoList)
}
