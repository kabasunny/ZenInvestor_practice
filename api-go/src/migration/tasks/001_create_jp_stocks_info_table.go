// api-go\src\migration\tasks\001_create_jp_stocks_info_table.go
package migration

import (
	"api-go/src/model"

	"gorm.io/gorm"
)

func Up_001(db *gorm.DB) error {
	return db.AutoMigrate(&model.JpStockInfo{})
}

func Down_001(db *gorm.DB) error {
	return db.Migrator().DropTable(&model.JpStockInfo{})
}
