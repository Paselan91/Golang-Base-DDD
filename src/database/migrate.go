package database

import (
	"gorm.io/gorm"
	"log"
	"my-module/config"
	"my-module/database/migrations"
)

func DBMigrate() (*gorm.DB, error) {
	log.Println("Migration start")
	db, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	migrations.MigrateUsersTable(db)
	migrations.MigrateContestsTable(db)
	migrations.MigratePlayersTable(db)

	log.Println("All migration has been successed")
	return db, nil
}
