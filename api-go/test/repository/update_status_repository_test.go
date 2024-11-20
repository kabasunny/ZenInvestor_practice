// api-go\test\repository\update_status_repository_test.go
package repository

import (
	"api-go/src/infra"
	"api-go/src/model"
	"api-go/src/repository"
	"fmt"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	godotenv.Load("../../.env") // 以前、テストではパスを指定しないとうまく読み取らなかった
	db := infra.SetupDB()
	return db
}

func cleanupTestData(db *gorm.DB) {
	// テストデータの削除
	db.Where("tb_name LIKE ?", "test_table%").Delete(&model.UpdateStatus{})
}

func backupAndClearTestData(db *gorm.DB) ([]model.UpdateStatus, error) {
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

func restoreTestData(db *gorm.DB, backup []model.UpdateStatus) {
	// バックアップしたデータを復元
	for _, data := range backup {
		db.Create(&data)
	}
}

// バックアップとクリアを除去した TestGetAllUpdateStatuses
func TestGetAllUpdateStatuses(t *testing.T) {
	db := setupTestDB()

	repo := repository.NewUpdateStatusRepository(db)
	statuses, err := repo.GetAllUpdateStatuses()

	assert.NoError(t, err)

	// 取得結果が空でないことを確認します
	assert.NotEmpty(t, statuses)

	// 結果をターミナルに表示（データ毎に改行）
	fmt.Println("GetAllUpdateStatuses:")
	for _, status := range statuses {
		fmt.Println(status)
	}
}

func TestUpdateStatus(t *testing.T) {
	db := setupTestDB()

	fmt.Println("TestUpdateStatus")

	// 既存のデータをバックアップしてクリア
	backup, err := backupAndClearTestData(db)
	assert.NoError(t, err)
	defer restoreTestData(db, backup)

	// テストデータの挿入
	yesterday := time.Now().AddDate(0, 0, -1)
	updateStatus := model.UpdateStatus{Date: yesterday, TbName: "test_table"}
	db.Create(&updateStatus)

	// 挿入するテストデータを表示
	fmt.Println("Before Update:", updateStatus)

	repo := repository.NewUpdateStatusRepository(db)
	err = repo.UpdateStatus("test_table")
	assert.NoError(t, err)

	// データベースから更新状態を取得して確認
	var result model.UpdateStatus
	db.First(&result, "tb_name = ?", "test_table")

	// 現在の日付
	today := time.Now().Format("2006-01-02")

	assert.Equal(t, "test_table", result.TbName)
	assert.Equal(t, today, result.Date.Format("2006-01-02"))

	// 更新後のデータの状態を表示
	fmt.Println("After Update:", result)

	// テストデータの削除
	cleanupTestData(db)
}

// テストの実行コード
// go test -v ./test/repository/update_status_repository_test.go -run TestGetAllUpdateStatuses TestUpdateStatus
// go test -v ./test/repository/update_status_repository_test.go -run TestUpdateStatus
// go test -v ./test/repository/update_status_repository_test.go -run TestGetAllUpdateStatuses
