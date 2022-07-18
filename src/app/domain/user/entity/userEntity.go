package entity

import (
	vo "my-module/app/domain/user/valueObject"
)

type UserEntity struct {
	id          uint
	userSubId   vo.UserSubId
	mailAddress vo.MailAddress
}

func NewUserEntity(
	id uint,
	userSubId vo.UserSubId,
	mailAddress vo.MailAddress,
) *UserEntity {
	return &UserEntity{
		id,
		userSubId,
		mailAddress,
	}
}

func (u *UserEntity) GetId() uint {
	return u.id
}

func (u *UserEntity) GetUserSubId() vo.UserSubId {
	return u.userSubId
}

func (u *UserEntity) GetMailAddress() vo.MailAddress {
	return u.mailAddress
}
