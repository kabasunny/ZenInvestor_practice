// api-go\src\migration\tasks\011_create_update_status_table.go
package migration

import (
	"api-go/src/model"

	"gorm.io/gorm"
)

func Up_011(db *gorm.DB) error {
	return db.AutoMigrate(&model.UpdateStatus{})
}

func Down_011(db *gorm.DB) error {
	return db.Migrator().DropTable(&model.UpdateStatus{})
}
