package contest

import (
	"github.com/labstack/echo"
	"my-module/app/infrastructure/repository"
	"my-module/app/usecases/contest/usecase"
	"my-module/config"
	"net/http"
)

type FetchAllContestsController struct{}

func NewFetchAllContestsController() *FetchAllContestsController {
	return &FetchAllContestsController{}
}

func (a *FetchAllContestsController) Execute(c echo.Context) error {
	db, err := config.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	r := repository.NewContestRepository(db)
	fetchAllContestsUseCase := usecase.NewFetchAllContestsUseCase(r)

	contestDtoList, errUc := fetchAllContestsUseCase.Execute()
	if errUc.HasErrors() {
		return c.JSON(http.StatusInternalServerError, errUc)
	}

	return c.JSON(http.StatusOK, contestDtoList)
}
