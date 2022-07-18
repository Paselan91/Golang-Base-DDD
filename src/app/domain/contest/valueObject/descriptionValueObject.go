package valueObject

import (
	"fmt"
	"my-module/app/errors"
	"unicode/utf8"
)

type Description struct {
	value string
}

const (
	JpnDescription = "説明"
	MaxDescription = 255
	MinDescription = 1
)

func NewDescription(description string) (*Description, *errors.AppError) {

	valueNum := utf8.RuneCountInString(description)
	if valueNum < MinDescription || valueNum > MaxDescription {
		msg := fmt.Sprintf("%sは%d文字以上%d文字以下の値にしてください\n", JpnDescription, MinDescription, MaxDescription)
		err := errors.NewAppError(msg)
		return nil, &err
	}
	return &Description{description}, nil
}

func (d Description) GetValue() string {
	return d.value
}

func (d Description) GetMax() int {
	return MaxDescription
}

func (d Description) GetMin() int {
	return MinDescription
}
