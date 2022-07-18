package routes

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"my-module/app/http/controllers/contest"
	"my-module/app/http/controllers/player"
	"my-module/app/http/controllers/user"
	"my-module/database"
	"my-module/database/seeds"
	"net/http"
)

func Run(e *echo.Echo, port string) {
	log.Printf("Server running at http://localhost:%s/", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func Routes(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello App !")
	})

	v1 := e.Group("/api/v1")

	// コンテスト一覧 TODO: ページネーション等で取得数絞る
	v1.GET("/contests", func(c echo.Context) error {
		fetchAllContestsController := contest.NewFetchAllContestsController()
		return fetchAllContestsController.Execute(c)
	})

	// コンテストに紐づくプレイヤー　TODO: ページネーション等で取得数絞る
	v1.GET("/contests/:contest_id/players", func(c echo.Context) error {
		fetchAllPlayersByContestIdController := player.NewFetchAllPlayersByContestIdController()
		return fetchAllPlayersByContestIdController.Execute(c)
	})

	// コンテストに紐づくランキング一覧　TODO: ページネーション等で取得数絞る
	v1.GET("/contests/:contest_id/ranking", func(c echo.Context) error {
		fetchRankingByContestIdController := player.NewFetchRankingByContestIdController()
		return fetchRankingByContestIdController.Execute(c)
	})

	// Test API
	// TODO: 別関数にしたい
	e.GET("/test_user/:id", func(c echo.Context) error {
		fetchUserByIdController := examples.NewFetchUserByIdController()
		return fetchUserByIdController.Execute(c)
	})
	e.POST("/test_user/create", func(c echo.Context) error {
		createUserController := examples.NewCreateUserController()
		return createUserController.Execute(c)
	})

	// Migration & Seed
	e.GET("/api/v1/migrate", migrate)
	e.GET("/api/v1/seed", seed)

}

// =============================
//    MIGRATE
// =============================
func migrate(c echo.Context) error {
	_, err := database.DBMigrate()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed migrate")
	} else {
		return c.String(http.StatusOK, "successed migrate")
	}
}

func seed(c echo.Context) error {
	_, err := seeds.Seeds()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed seed")
	} else {
		return c.String(http.StatusOK, "successed seed")
	}
}
