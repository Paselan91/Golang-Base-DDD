package valueObject

import (
	"fmt"
	// "log"
	"my-module/app/errors"
)

type SortType struct {
	value SortEnum
}

const (
	JpnSortType = "ソートタイプ"
)

type SortEnum string

const (
	Desc      = SortEnum("Desc")      // 作成日の降順 //TODO: 共通SortTypeVOを継承できないか
	ScoreDesc = SortEnum("ScoreDesc") // players.scoreの降順
)

func types() [2]SortEnum {
	types := [2]SortEnum{
		Desc,
		ScoreDesc,
	}

	return types
}

func NewSortType(sortType SortEnum) (*SortType, *errors.AppError) {

	ok := isTypes(sortType)
	if !ok {
		msg := fmt.Sprintf("%sへ不正な値が指定されました。sortType: %s\n", JpnSortType, sortType)
		err := errors.NewAppError(msg)
		return nil, &err
	}
	return &SortType{sortType}, nil
}

func (s SortType) GetValue() SortEnum {
	return s.value
}

func isTypes(value SortEnum) bool {
	for _, sortEnum := range types() {
		if value == sortEnum {
			return true
		}
	}
	return false
}
