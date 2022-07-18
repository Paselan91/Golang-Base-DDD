package seeds

import (
	"fmt"
	"my-module/app/infrastructure/models"
	"my-module/config"
	"strings"
)

func ContestsTableSeeder() (bool, error) {
	db, err := config.ConnectDB()
	if err != nil {
		return false, err
	}

	// 生成するデータ数
	const testNum = 50

	// ユーザーを取得
	users := []models.User{}
	db.Find(&users)

	contests := []models.Contest{}
	contest := models.Contest{}

	// description のデータ（長い文字列を生成）
	description := strings.Repeat("これはコンテストの説明テストです。", 10)

	for i := 0; i < testNum; i++ {
		contest = models.Contest{
			Name:        fmt.Sprintf("XXXXそうなコンテスト_%d", i),
			Description: fmt.Sprintf("%d_%s", i, description),
			ImagePath:   fmt.Sprintf("https://www.test_%d.com", i),
			User:        users[i],
		}
		contests = append(contests, contest)
	}

	if err := db.Create(&contests).Error; err != nil {
		return false, err
	}
	return true, nil
}
