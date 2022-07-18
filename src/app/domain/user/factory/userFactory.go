package factory

import (
	"my-module/app/domain/user/entity"
	vo "my-module/app/domain/user/valueObject"
	"my-module/app/errors"
)

type UserFactory struct{}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func (uf *UserFactory) Create(
	userSubId string,
	mailAddress string,
) (*entity.UserEntity, *errors.AppError) {
	userSubIdVo, err := vo.NewUserSubId(userSubId)
	if err.HasErrors() {
		return nil, err
	}

	mailAddressVo, err := vo.NewMailAddress(mailAddress)
	if err.HasErrors() {
		return nil, err
	}

	userEntity := entity.NewUserEntity(
		0,
		*userSubIdVo,
		*mailAddressVo,
	)
	return userEntity, nil
}

func (uf *UserFactory) Reconstruct(
	id uint,
	userSubId vo.UserSubId,
	mailAddress vo.MailAddress,
) (*entity.UserEntity, *errors.AppError) {

	userEntity := entity.NewUserEntity(
		id,
		userSubId,
		mailAddress,
	)
	return userEntity, nil
}
