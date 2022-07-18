package seeds

import (
	"gorm.io/gorm"
	"log"
	"my-module/config"
)

func Seeds() (*gorm.DB, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}

	UsersTableSeeder()
	ContestsTableSeeder()
	PlayersTableSeeder()

	log.Println("Seeded All Seeders")
	return conn, nil

}
