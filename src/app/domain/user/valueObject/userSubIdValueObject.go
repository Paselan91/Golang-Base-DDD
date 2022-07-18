package valueObject

import (
	"fmt"
	"my-module/app/errors"
	"unicode/utf8"
)

type UserSubId struct {
	value string
}

const (
	JpnUserSubId = "ユーザーID"
	MaxUserSubId = 255
	MinUserSubId = 1
)

func NewUserSubId(userSubId string) (*UserSubId, *errors.AppError) {

	if utf8.RuneCountInString(userSubId) < 0 || utf8.RuneCountInString(userSubId) > MaxUserSubId {
		msg := fmt.Sprintf("%sは%d文字以上%d以上の値にしてください\n", JpnUserSubId, MinUserSubId, MaxUserSubId)
		err := errors.NewAppError(msg)
		return nil, &err
	}
	return &UserSubId{userSubId}, nil
}

func (u UserSubId) GetValue() string {
	return u.value
}

func (u UserSubId) GetMax() int {
	return MaxUserSubId
}

func (u UserSubId) GetMin() int {
	return MinUserSubId
}
