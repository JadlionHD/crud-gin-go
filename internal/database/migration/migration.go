package migration

import "gorm.io/gorm"

func MigrateTable(db *gorm.DB) {
	db.AutoMigrate(&PostDB{})
}
