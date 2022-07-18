package migrations

import (
	"gorm.io/gorm"
	"my-module/app/infrastructure/models"
)

func MigratePlayersTable(db *gorm.DB) {
	if db.Migrator().HasTable(models.Player{}) {
		db.Migrator().DropTable(models.Player{})
	}
	db.AutoMigrate(models.Player{})
}
