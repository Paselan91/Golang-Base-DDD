package player

import (
	"github.com/labstack/echo"
	"my-module/app/infrastructure/repository"
	"my-module/app/usecases/player/usecase"
	"my-module/config"
	"net/http"
	"strconv"
)

type FetchAllPlayersByContestIdController struct{}

func NewFetchAllPlayersByContestIdController() *FetchAllPlayersByContestIdController {
	return &FetchAllPlayersByContestIdController{}
}

func (controller *FetchAllPlayersByContestIdController) Execute(c echo.Context) error {
	contestId, err := strconv.Atoi(c.Param("contest_id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	db, err := config.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	r := repository.NewPlayerRepository(db)
	fetchAllPlayersByContestIdUseCase := usecase.NewFetchAllPlayersByContestIdUseCase(r)

	playerDtoList, errUc := fetchAllPlayersByContestIdUseCase.Execute(contestId)
	if errUc.HasErrors() {
		return c.JSON(http.StatusInternalServerError, errUc)
	}

	return c.JSON(http.StatusOK, playerDtoList)
}
