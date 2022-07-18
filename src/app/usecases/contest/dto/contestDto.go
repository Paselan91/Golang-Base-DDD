package dto

import ()

type ContestDto struct {
	Id          uint   `json:"id"`
	UserSubId   string `json:"user_sub_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImagePath   string `json:"image_path"`
}

func NewContestDto(id uint, userSubId, name, description, imagePath string) *ContestDto {
	return &ContestDto{id, userSubId, name, description, imagePath}
}
