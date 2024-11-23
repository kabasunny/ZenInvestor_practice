// api-go\test\repository\repositry_test_helper\repository_test_helper.go

package repository_test_helper

import (
	"api-go/src/infra"
	"api-go/src/model"
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SetupTestDB() *gorm.DB {
	godotenv.Load("../../.env") // テストではパスを指定しないとうまく読み取らなかった
	db := infra.SetupDB()
	return db
}

func InitializeUpdateStatusTable(db *gorm.DB) {
	initialRecords := []model.UpdateStatus{
		{TbName: "jp_5d_mva_ranking", Date: time.Now().AddDate(0, 0, -1)}, // 1日前
		{TbName: "jp_daily_price", Date: time.Now().AddDate(0, 0, -1)},    // 1日前
		{TbName: "jp_stocks_info", Date: time.Now()},                      // 本日の日付に修正
	}
	for _, record := range initialRecords {
		db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&record)
	}
}

func PrintUpdateStatusTable(db *gorm.DB) {
	var statuses []model.UpdateStatus
	db.Find(&statuses)
	for _, status := range statuses {
		fmt.Printf("Table: %s, Date: %s\n", status.TbName, status.Date)
	}
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
