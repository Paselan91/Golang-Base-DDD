package examples

import (
	"github.com/labstack/echo"
	"my-module/app/http/inputs/userInputs"
	"my-module/app/infrastructure/repository"
	"my-module/app/usecases/user/usecase"
	"my-module/config"
	"net/http"
	// "strconv"
)

type CreateUserController struct{}

func NewCreateUserController() *CreateUserController {
	return &CreateUserController{}
}

func (a *CreateUserController) Execute(c echo.Context) error {

	createUserInput := userInputs.GetCreateUserInput(c)

	db, err := config.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	ur := repository.NewUserRepository(db)
	createUserUseCase := usecase.NewCreateUserUseCase(ur)

	user, errUc := createUserUseCase.Execute(createUserInput)
	if errUc.HasErrors() {
		return c.JSON(http.StatusInternalServerError, errUc)
	}

	return c.JSON(http.StatusOK, user)
}
