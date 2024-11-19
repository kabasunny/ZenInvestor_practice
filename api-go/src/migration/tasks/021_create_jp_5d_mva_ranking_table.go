// api-go\src\migration\tasks\021_create_jp_5d_mva_ranking_table.go
package migration

import (
	"api-go/src/model"

	"gorm.io/gorm"
)

func Up_021(db *gorm.DB) error {
	return db.AutoMigrate(&model.Jp5dMvaRanking{})
}

func Down_021(db *gorm.DB) error {
	return db.Migrator().DropTable(&model.Jp5dMvaRanking{})
}
