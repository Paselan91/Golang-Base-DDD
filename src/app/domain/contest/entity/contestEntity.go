package entity

import (
	vo "my-module/app/domain/contest/valueObject"
	"my-module/app/domain/user/entity"
)

type ContestEntity struct {
	id          uint
	userEntity  *entity.UserEntity
	name        vo.Name
	description vo.Description
	imagePath   vo.ImagePath
}

func NewContestEntity(id uint, userEntity *entity.UserEntity, name vo.Name, description vo.Description, imagePath vo.ImagePath) *ContestEntity {
	return &ContestEntity{
		id:          id,
		userEntity:  userEntity,
		name:        name,
		description: description,
		imagePath:   imagePath,
	}
}

func (c *ContestEntity) GetId() uint {
	return c.id
}

func (c *ContestEntity) GetUserEntity() *entity.UserEntity {
	return c.userEntity
}

func (c *ContestEntity) GetName() vo.Name {
	return c.name
}

func (c *ContestEntity) GetDescription() vo.Description {
	return c.description
}

func (c *ContestEntity) GetImagePath() vo.ImagePath {
	return c.imagePath
}
