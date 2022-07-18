package examples

import (
	"github.com/labstack/echo"
	"my-module/app/infrastructure/repository"
	"my-module/app/usecases/user/usecase"
	"my-module/config"
	"net/http"
	"strconv"
)

type DeleteUserController struct{}

func NewDeleteUserController() *DeleteUserController {
	return &DeleteUserController{}
}

func (a *DeleteUserController) Execute(c echo.Context) error {

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	db, err := config.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	ur := repository.NewUserRepository(db)
	deleteUserUseCase := usecase.NewDeleteUserUseCase(ur)

	deletedUserId, errUc := deleteUserUseCase.Execute(userId)
	if errUc.HasErrors() {
		return c.JSON(http.StatusInternalServerError, errUc)
	}

	return c.JSON(http.StatusOK, deletedUserId)
}
