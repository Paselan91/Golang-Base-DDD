package repository

import (
	"gorm.io/gorm"
	"log"
	"my-module/app/domain/user/entity"
	"my-module/app/domain/user/repository"
	"my-module/app/errors"
	"my-module/app/infrastructure/models"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindById(userId int) (*entity.UserEntity, *errors.AppError) {

	user := models.User{}

	err := ur.db.Find(&user, userId).Error
	if err != nil {
		err := errors.NewAppError("ユーザーが見つかりませんでした。")
		return nil, &err
	}

	ue, errUe := user.ToEntity()
	if errUe.HasErrors() {
		return nil, errUe
	}

	return ue, nil
}

func (ur *userRepository) Create(userEntity *entity.UserEntity) (*entity.UserEntity, *errors.AppError) {

	user := models.User{
		UserSubId:   userEntity.GetUserSubId().GetValue(),
		MailAddress: userEntity.GetMailAddress().GetValue(),
	}

	result := ur.db.Create(&user)
	err := result.Error
	if err != nil {
		err := errors.NewAppError("ユーザーの新規登録に失敗しました。")
		return nil, &err
	}

	// TODO: userにすべて作成した値が入っているか確認する
	log.Println("user")
	log.Println(user)

	ue, errUe := user.ToEntity()
	if errUe.HasErrors() {
		return nil, errUe
	}

	return ue, nil
}

func (ur *userRepository) Delete(userId int) (bool, *errors.AppError) {

	user := models.User{}

	err := ur.db.Delete(&user, userId).Error
	if err != nil {
		err := errors.NewAppError("ユーザーの削除に失敗しました。")
		return false, &err
	}

	return true, nil
}
