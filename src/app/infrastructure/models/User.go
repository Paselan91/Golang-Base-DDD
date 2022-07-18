package models

import (
	"my-module/app/domain/user/entity"
	vo "my-module/app/domain/user/valueObject"
	"my-module/app/errors"
	// "log"
)

type User struct {
	Model
	UserSubId   string `gorm:"not null;unique"`
	MailAddress string `gorm:"not null;unique"`
}

func (u *User) ToEntity() (*entity.UserEntity, *errors.AppError) {

	userSubIdVo, err := vo.NewUserSubId(u.UserSubId)
	if err.HasErrors() {
		return nil, err
	}

	mailAddressVo, err := vo.NewMailAddress(u.MailAddress)
	if err.HasErrors() {
		return nil, err
	}

	entity := entity.NewUserEntity(
		u.Id,
		*userSubIdVo,
		*mailAddressVo,
	)

	return entity, nil
}
