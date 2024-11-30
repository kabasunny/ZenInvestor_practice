// api-go\test\repository\update_status_repository_test.go
package repository

import (
	"api-go/src/model"
	"api-go/src/repository"
	"api-go/test/repository/repository_test_helper"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAllUpdateStatuses(t *testing.T) {
	db := repository_test_helper.SetupTestDB()

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
	db := repository_test_helper.SetupTestDB()

	fmt.Println("TestUpdateStatus")

	// 既存のデータをバックアップしてクリア
	backup, err := repository_test_helper.BackupAndClearTestData(db)
	assert.NoError(t, err)
	defer repository_test_helper.RestoreTestData(db, backup)

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
	repository_test_helper.CleanupUpdateStatusTestData(db)
}

func TestUpsertStatus(t *testing.T) {
	db := repository_test_helper.SetupTestDB()

	fmt.Println("TestUpsertStatus")

	// テスト環境のクリーンアップ
	repository_test_helper.CleanupUpdateStatusTestData(db)

	repo := repository.NewUpdateStatusRepository(db)

	// 新しいテーブル名でアップサート
	err := repo.UpdateStatus("new_table")
	assert.NoError(t, err)

	// 新しいテーブル名のレコードが存在することを確認
	var newResult model.UpdateStatus
	db.First(&newResult, "tb_name = ?", "new_table")
	assert.Equal(t, "new_table", newResult.TbName)
	assert.Equal(t, time.Now().Format("2006-01-02"), newResult.Date.Format("2006-01-02"))

	// 既存のテーブル名でアップサート（更新）
	yesterday := time.Now().AddDate(0, 0, -1)
	existingStatus := model.UpdateStatus{Date: yesterday, TbName: "existing_table"}
	db.Create(&existingStatus)

	// 挿入するテストデータを表示
	fmt.Println("Before Update:", existingStatus)

	err = repo.UpdateStatus("existing_table")
	assert.NoError(t, err)

	// 既存のテーブル名のレコードが更新されたことを確認
	var existingResult model.UpdateStatus
	db.First(&existingResult, "tb_name = ?", "existing_table")
	assert.Equal(t, "existing_table", existingResult.TbName)
	assert.Equal(t, time.Now().Format("2006-01-02"), existingResult.Date.Format("2006-01-02"))

	// 更新後のデータの状態を表示
	fmt.Println("After Update:", existingResult)

	// テストデータの削除
	repository_test_helper.CleanupUpdateStatusTestData(db)
}

// go test -v ./test/repository/update_status_repository_test.go -run TestUpdateStatus
// go test -v ./test/repository/update_status_repository_test.go -run TestGetAllUpdateStatuses
// go test -v ./test/repository/update_status_repository_test.go -run TestUpsertStatus
