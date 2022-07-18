package dto

import ()

type PlayerDto struct {
	Id          uint   `json:"id"`
	UserSubId   string `json:"user_sub_id"`
	ContestId   uint   `json:"contest_id"`
	Name        string `json:"name"`
	ImagePath   string `json:"image_path"`
	Description string `json:"description"`
	Score       int    `json:"score"`
	Clicked     uint   `json:"clicked"`
}

func NewPlayerDto(
	id uint,
	userSubId string,
	contestId uint,
	name string,
	imagePath string,
	description string,
	score int,
	clicked uint,
) *PlayerDto {
	return &PlayerDto{
		id,
		userSubId,
		contestId,
		name,
		imagePath,
		description,
		score,
		clicked,
	}
}
