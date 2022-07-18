package userInputs

import (
	"github.com/labstack/echo"
	"log"
)

type CreateUserInput struct {
	UserSubId   string `json:"user_sub_id"`
	MailAddress string `json:"mail_address"`
}

// TODO: カスタムエラーを返す
func GetCreateUserInput(c echo.Context) *CreateUserInput {

	createUserInput := new(CreateUserInput)
	if err := c.Bind(createUserInput); err != nil {
		log.Println("Error!") // TODO: カスタムエラーを返す
	}

	return createUserInput
}

func (i *CreateUserInput) GetUserSubId() string {
	return i.UserSubId
}

func (i *CreateUserInput) GetMailAddress() string {
	return i.MailAddress
}
