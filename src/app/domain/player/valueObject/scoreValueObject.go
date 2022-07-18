package valueObject

import (
	"fmt"
	"my-module/app/errors"
	// "log"
)

type Score struct {
	value int
}

const (
	JpnNameScore = "点数"
	MaxScore     = 1000000 // FIXME: 暫定値
	MinScore     = 0
)

func NewScore(score int) (*Score, *errors.AppError) {

	if score < MinScore || score > MaxScore {
		msg := fmt.Sprintf("%sは%d文字以上%d以下の値にしてください\n", JpnNameScore, MinScore, MaxScore)
		err := errors.NewAppError(msg)
		return nil, &err
	}

	return &Score{score}, nil
}

func (s Score) GetValue() int {
	return s.value
}

func (s Score) GetMax() int {
	return MaxScore
}

func (s Score) GetMin() int {
	return MinScore
}
