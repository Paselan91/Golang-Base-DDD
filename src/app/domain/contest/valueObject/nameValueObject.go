package valueObject

import (
	"fmt"
	"my-module/app/errors"
	"unicode/utf8"
)

type Name struct {
	value string
}

const (
	JpnName = "名称"
	MaxName = 20
	MinName = 1
)

func NewName(name string) (*Name, *errors.AppError) {

	valueNum := utf8.RuneCountInString(name)
	if valueNum < MinName || valueNum > MaxName {
		msg := fmt.Sprintf("%sは%d文字以上%d文字以下の値にしてください\n", JpnName, MinName, MaxName)
		err := errors.NewAppError(msg)
		return nil, &err
	}
	return &Name{name}, nil
}

func (n Name) GetValue() string {
	return n.value
}

func (n Name) GetMax() int {
	return MaxName
}

func (n Name) GetMin() int {
	return MinName
}
