package seeds

import (
	"fmt"
	"math/rand"
	vo "my-module/app/domain/player/valueObject"
	"my-module/app/infrastructure/models"
	"my-module/config"
	"strings"
)

func PlayersTableSeeder() (bool, error) {

	// 生成するデータ数
	// 1コンテストあたりの参加者数
	const testPlayersPerContest = 10
	// 紐づけるコンテストの数
	const testContests = 10

	score := vo.Score{}

	mixScore := score.GetMin()
	maxScore := score.GetMax()

	db, err := config.ConnectDB()
	if err != nil {
		return false, err
	}

	// ユーザーを取得
	users := []models.User{}
	db.Find(&users)

	// コンテストを取得
	contests := []models.Contest{}
	db.Find(&contests)

	players := []models.Player{}
	player := models.Player{}

	// description のデータ（長い文字列を生成）
	description := strings.Repeat("これは参加者の説明です。", 20)

	// testPlayersPerContest * testContests の数だけテストデータ作成
	k := 0
	for i := 0; i < testContests; i++ {
		for j := 0; j < testPlayersPerContest; j++ {

			player = models.Player{
				User:        users[k],
				Contest:     contests[i],
				Name:        fmt.Sprintf("参加者_%d", k),
				ImagePath:   fmt.Sprintf("https://www.test_%d.com", k),
				Description: fmt.Sprintf("%d_%s", k, description),
				Score:       rand.Intn(maxScore-mixScore) + mixScore,
				Clicked:     uint(rand.Intn(1000)),
			}
			players = append(players, player)
			k++
		}
	}

	if err := db.Create(&players).Error; err != nil {
		return false, err
	}
	return true, nil
}
