// api-go\test\repository\repositry_test_helper\repository_test_helper.go

package repository_test_helper

import (
	"api-go/src/infra"
	"api-go/src/model"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	godotenv.Load("../../.env") // テストではパスを指定しないとうまく読み取らなかった
	db := infra.SetupDB()
	return db
}

func CleanupStockTestData(db *gorm.DB) {
	// テストデータの削除
	db.Where("ticker LIKE ?", "test%").Delete(&model.JpStockInfo{})
}

func CleanupUpdateStatusTestData(db *gorm.DB) {
	// テストデータの削除
	db.Where("tb_name LIKE ?", "test_table%").Delete(&model.UpdateStatus{})
}

func BackupAndClearTestData(db *gorm.DB) ([]model.UpdateStatus, error) {
	var backup []model.UpdateStatus
	// 既存のデータをバックアップ
	if err := db.Find(&backup).Error; err != nil {
		return nil, err
	}
	// テストデータをクリア
	if err := db.Exec("DELETE FROM update_status").Error; err != nil {
		return nil, err
	}
	return backup, nil
}

func RestoreTestData(db *gorm.DB, backup []model.UpdateStatus) {
	// バックアップしたデータを復元
	for _, data := range backup {
		db.Create(&data)
	}
}
