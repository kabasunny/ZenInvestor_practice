// api-go\test\repository\update_status_repository_test.go

package repository

import (
	"api-go/src/infra"
	"api-go/src/model"
	"api-go/src/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	// データベース接続の初期化
	db := infra.SetupDB()
	return db
}

func cleanupTestData(db *gorm.DB) {
	// テストデータの削除
	db.Where("tb_name LIKE ?", "test_table_%").Delete(&model.UpdateStatus{})
}

func TestGetAllUpdateStatuses(t *testing.T) {
	db := setupTestDB()

	// テストデータの挿入
	yesterday := time.Now().AddDate(0, 0, -1)
	updateStatuses := []model.UpdateStatus{
		{Date: yesterday, TbName: "test_table_1"},
		{Date: yesterday, TbName: "test_table_2"},
	}
	for _, status := range updateStatuses {
		db.Create(&status)
	}

	repo := repository.NewUpdateStatusRepository(db)
	statuses, err := repo.GetAllUpdateStatuses()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(statuses))
	assert.Equal(t, "test_table_1", statuses[0].TbName)
	assert.Equal(t, "test_table_2", statuses[1].TbName)

	// テストデータの削除
	cleanupTestData(db)
}

func TestUpdateStatus(t *testing.T) {
	db := setupTestDB()

	// テストデータの挿入
	yesterday := time.Now().AddDate(0, 0, -1)
	updateStatus := model.UpdateStatus{Date: yesterday, TbName: "test_table"}

	repo := repository.NewUpdateStatusRepository(db)
	err := repo.UpdateStatus(&updateStatus)

	assert.NoError(t, err)

	// データベースから更新状態を取得して確認
	var result model.UpdateStatus
	db.First(&result, "tb_name = ?", "test_table")
	assert.Equal(t, "test_table", result.TbName)
	assert.Equal(t, yesterday.Format("2006-01-02"), result.Date.Format("2006-01-02"))

	// テストデータの削除
	cleanupTestData(db)
}

// テストの実行コード
// go test -v ./test/repository/update_status_repository_test.go -run TestGetAllUpdateStatuses TestUpdateStatus
