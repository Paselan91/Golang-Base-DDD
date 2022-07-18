package migrations

import (
	"gorm.io/gorm"
	"my-module/app/infrastructure/models"
)

func MigrateUsersTable(db *gorm.DB) {
	if db.Migrator().HasTable(models.User{}) {
		db.Migrator().DropTable(models.User{})
	}
	db.AutoMigrate(models.User{})
}
