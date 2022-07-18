package models

import (
	cEntity "my-module/app/domain/contest/entity"
	pEntity "my-module/app/domain/player/entity"
	vo "my-module/app/domain/player/valueObject"
	uEntity "my-module/app/domain/user/entity"
	"my-module/app/errors"
	// "log"
)

type Player struct {
	Model
	UserId      uint `gorm:"not null"`
	User        User
	ContestId   uint `gorm:"not null"`
	Contest     Contest
	Name        string `gorm:"not null;unique"`
	ImagePath   string `gorm:"not null"`
	Description string `gorm:"not null"`
	Score       int    `gorm:"not null"`
	Clicked     uint   `gorm:"not null"`
}

func (player *Player) ToEntity() (*pEntity.PlayerEntity, *errors.AppError) {

	nameVo, err := vo.NewName(player.Name)
	if err.HasErrors() {
		return nil, err
	}

	imagePathVo, err := vo.NewImagePath(player.ImagePath)
	if err.HasErrors() {
		return nil, err
	}

	descriptionVo, err := vo.NewDescription(player.Description)
	if err.HasErrors() {
		return nil, err
	}

	scoreVo, err := vo.NewScore(player.Score)
	if err.HasErrors() {
		return nil, err
	}

	clickedVo, err := vo.NewClicked(player.Clicked)
	if err.HasErrors() {
		return nil, err
	}

	userEntity := &uEntity.UserEntity{}
	if player.User.Id != 0 {
		userEntity, err = player.User.ToEntity()
		if err.HasErrors() {
			return nil, err
		}
	}

	contestEntity := &cEntity.ContestEntity{}
	if player.Contest.Id != 0 {
		contestEntity, err = player.Contest.ToEntity()
		if err.HasErrors() {
			return nil, err
		}
	}

	entity := pEntity.NewPlayerEntity(
		player.Id,
		userEntity,
		contestEntity,
		*nameVo,
		*imagePathVo,
		*descriptionVo,
		*scoreVo,
		*clickedVo,
	)

	return entity, nil
}
