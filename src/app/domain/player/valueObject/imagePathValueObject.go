package valueObject

import (
	"fmt"
	"my-module/app/errors"
	"regexp"
	"unicode/utf8"
)

type ImagePath struct {
	value string
}

const (
	JpnImagePath = "画像パス"
	MaxImagePath = 255
	MinImagePath = 1
)

// TODO: 共通化する BaseUrl ?
func NewImagePath(imagePath string) (*ImagePath, *errors.AppError) {

	valueNum := utf8.RuneCountInString(imagePath)
	if valueNum < MinImagePath || valueNum > MaxImagePath {
		msg := fmt.Sprintf("%sは%d文字以上%d文字以下の値にしてください\n", JpnImagePath, MinImagePath, MaxImagePath)
		err := errors.NewAppError(msg)
		return nil, &err
	}
	return &ImagePath{imagePath}, nil
}

func (i ImagePath) GetValue() string {
	return i.value
}

func (i ImagePath) GetMax() int {
	return MaxImagePath
}

func (i ImagePath) GetMin() int {
	return MinImagePath
}

func isUrl(url string) (bool, *errors.AppError) {
	const regex = `https?://[\w!\?/\+\-_~=;\.,\*&@#\$%\(\)'\[\]]+`

	match, err := regexp.MatchString(regex, url)
	if err != nil {
		err := errors.NewAppError("regexpに失敗しました。")
		return false, &err
	}
	if !match {
		return false, nil
	}
	return true, nil
}
