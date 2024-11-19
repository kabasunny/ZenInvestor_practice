// api-go\src\repository\update_status_repository.go
package repository

import (
	"api-go/src/model"
)

type UpdateStatusRepository interface {
	// 全ての更新状態をチェックする: データ取得時(毎回)
	GetAllUpdateStatuses() ([]model.UpdateStatus, error)

	// 更新状態を更新する: テーブル更新時(毎回)
	UpdateStatus(updateStatus *model.UpdateStatus) error
}
