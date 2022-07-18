package valueObject

import (
	"fmt"
	"my-module/app/errors"
)

type Clicked struct {
	value uint
}

const (
	JpnClicked = "クリックされた数"
	MaxClicked = 1000000
	MinClicked = 0
)

func NewClicked(clicked uint) (*Clicked, *errors.AppError) {

	if clicked > MaxClicked {
		msg := fmt.Sprintf("%sは%d以下の値にしてください\n", JpnClicked, MaxClicked)
		err := errors.NewAppError(msg)
		return nil, &err
	}
	return &Clicked{clicked}, nil
}

func (c Clicked) GetValue() uint {
	return c.value
}

func (c Clicked) GetMax() int {
	return MaxClicked
}

func (c Clicked) GetMin() int {
	return MinClicked
}
