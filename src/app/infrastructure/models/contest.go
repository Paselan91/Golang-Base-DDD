package models

import (
	cEntity "my-module/app/domain/contest/entity"
	vo "my-module/app/domain/contest/valueObject"
	uEntity "my-module/app/domain/user/entity"
	"my-module/app/errors"
	// "log"
)

type Contest struct {
	Model
	UserId      uint `gorm:"not null"`
	User        User
	Name        string `gorm:"not null;unique"`
	Description string `gorm:"not null"`
	ImagePath   string `gorm:"not null"`
}

func (contest *Contest) ToEntity() (*cEntity.ContestEntity, *errors.AppError) {

	nameVo, err := vo.NewName(contest.Name)
	if err.HasErrors() {
		return nil, err
	}

	descriptionVo, err := vo.NewDescription(contest.Description)
	if err.HasErrors() {
		return nil, err
	}

	imagePathVo, err := vo.NewImagePath(contest.ImagePath)
	if err.HasErrors() {
		return nil, err
	}

	userEntity := &uEntity.UserEntity{}
	if contest.User.Id != 0 {
		userEntity, err = contest.User.ToEntity()
		if err.HasErrors() {
			return nil, err
		}
	}

	entity := cEntity.NewContestEntity(contest.Id, userEntity, *nameVo, *descriptionVo, *imagePathVo)

	return entity, nil
}
