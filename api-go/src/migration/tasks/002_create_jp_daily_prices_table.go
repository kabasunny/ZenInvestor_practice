// api-go\src\migration\tasks\002_create_jp_daily_prices_table.go
package migration

import (
	"api-go/src/model"

	"gorm.io/gorm"
)

func Up_002(db *gorm.DB) error {
	return db.AutoMigrate(&model.JpDailyPrice{})
}

func Down_002(db *gorm.DB) error {
	return db.Migrator().DropTable(&model.JpDailyPrice{})
}
