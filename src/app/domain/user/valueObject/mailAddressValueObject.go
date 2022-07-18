package valueObject

import (
	"fmt"
	"log"
	"my-module/app/errors"
	"regexp"
	"unicode/utf8"
)

type MailAddress struct {
	value string
}

const (
	JpnMailAddress = "メールアドレス"
	MaxMailAddress = 255
	MinMailAddress = 1
)

func NewMailAddress(mailAddress string) (*MailAddress, *errors.AppError) {

	if utf8.RuneCountInString(mailAddress) < 0 && utf8.RuneCountInString(mailAddress) > MaxMailAddress {
		msg := fmt.Sprintf("%sは%d文字以上%d以下の値にしてください\n", JpnMailAddress, MinMailAddress, MaxMailAddress)
		err := errors.NewAppError(msg)
		return nil, &err
	}

	ok, err := isMailAddress(mailAddress)
	if err.HasErrors() {
		return nil, err
	}
	if !ok {
		msg := fmt.Sprintf("%sは%s形式ではありません。", mailAddress, JpnMailAddress)
		log.Println(msg)
		log.Println(mailAddress)
		err := errors.NewAppError(msg)
		return nil, &err
	}

	return &MailAddress{mailAddress}, nil
}

func (m MailAddress) GetValue() string {
	return m.value
}

func (m MailAddress) GetMax() int {
	return MaxMailAddress
}

func (m MailAddress) GetMin() int {
	return MinMailAddress
}

func isMailAddress(mailAddress string) (bool, *errors.AppError) {
	const regex = `^(?i:[^ @"<>]+|".*")@(?i:[a-z1-9.])+.(?i:[a-z])+$`

	match, err := regexp.MatchString(regex, mailAddress)
	if err != nil {
		err := errors.NewAppError("regexpに失敗しました。")
		return false, &err
	}
	if !match {
		return false, nil
	}
	return true, nil
}
