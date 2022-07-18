package migrations

import (
	"gorm.io/gorm"
	"my-module/app/infrastructure/models"
)

func MigrateContestsTable(db *gorm.DB) {
	if db.Migrator().HasTable(models.Contest{}) {
		db.Migrator().DropTable(models.Contest{})
	}
	db.AutoMigrate(models.Contest{})
}
