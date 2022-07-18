package entity

import (
	cEntity "my-module/app/domain/contest/entity"
	vo "my-module/app/domain/player/valueObject"
	uEntity "my-module/app/domain/user/entity"
)

type PlayerEntity struct {
	id            uint
	userEntity    *uEntity.UserEntity
	contestEntity *cEntity.ContestEntity
	name          vo.Name
	imagePath     vo.ImagePath
	description   vo.Description
	score         vo.Score
	clicked       vo.Clicked
}

func NewPlayerEntity(
	id uint,
	userEntity *uEntity.UserEntity,
	contestEntity *cEntity.ContestEntity,
	name vo.Name,
	imagePath vo.ImagePath,
	description vo.Description,
	score vo.Score,
	clicked vo.Clicked,
) *PlayerEntity {
	return &PlayerEntity{
		id,
		userEntity,
		contestEntity,
		name,
		imagePath,
		description,
		score,
		clicked,
	}
}

func (c *PlayerEntity) GetId() uint {
	return c.id
}

func (c *PlayerEntity) GetUserEntity() *uEntity.UserEntity {
	return c.userEntity
}

func (c *PlayerEntity) GetContestEntity() *cEntity.ContestEntity {
	return c.contestEntity
}

func (c *PlayerEntity) GetName() vo.Name {
	return c.name
}

func (c *PlayerEntity) GetImagePath() vo.ImagePath {
	return c.imagePath
}

func (c *PlayerEntity) GetDescription() vo.Description {
	return c.description
}

func (c *PlayerEntity) GetScore() vo.Score {
	return c.score
}

func (c *PlayerEntity) GetClicked() vo.Clicked {
	return c.clicked
}
