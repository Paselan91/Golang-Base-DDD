package dto

import ()

type UserDto struct {
	Id          uint   `json:"id"`
	UserSubId   string `json:"user_sub_id"`
	MailAddress string `json:"mail_address"`
}

func NewUserDto(
	id uint,
	userSubId string,
	mailAddress string,
) *UserDto {
	return &UserDto{
		id,
		userSubId,
		mailAddress,
	}
}
