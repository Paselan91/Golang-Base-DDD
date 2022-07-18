package examples

import (
	"github.com/labstack/echo"
	"my-module/app/infrastructure/repository"
	"my-module/app/usecases/user/usecase"
	"my-module/config"
	"net/http"
	"strconv"
)

type FetchUserByIdController struct{}

func NewFetchUserByIdController() *FetchUserByIdController {
	return &FetchUserByIdController{}
}

func (a *FetchUserByIdController) Execute(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	db, err := config.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	ur := repository.NewUserRepository(db)
	fetchUserUseCase := usecase.NewFetchUserUseCase(ur)

	user, errUc := fetchUserUseCase.Execute(userId)
	if errUc.HasErrors() {
		return c.JSON(http.StatusInternalServerError, errUc)
	}

	return c.JSON(http.StatusOK, user)
}
