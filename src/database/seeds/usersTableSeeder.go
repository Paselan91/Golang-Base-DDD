package seeds

import (
	"fmt"
	"my-module/app/infrastructure/models"
	"my-module/config"
)

func UsersTableSeeder() (bool, error) {
	db, err := config.ConnectDB()
	if err != nil {
		return false, err
	}

	// 生成するデータ数
	const testNum = 120

	users := []models.User{}
	user := models.User{}
	for i := 0; i < testNum; i++ {
		user = models.User{
			UserSubId:   fmt.Sprintf("test_user_%d", i),
			MailAddress: fmt.Sprintf("test_user_%d@test.com", i),
		}
		users = append(users, user)
	}

	if err := db.Create(&users).Error; err != nil {
		return false, err
	}
	return true, nil
}
