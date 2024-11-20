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

// テストの実行コード
// go test -v ./test/repository/update_status_repository_test.go -run TestUpdateStatus
// go test -v ./test/repository/update_status_repository_test.go -run TestGetAllUpdateStatuses
