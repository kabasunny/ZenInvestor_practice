// api-go\src\repository\update_status_repository.go

package repository

type UpdateStatusRepository interface {
	// 更新状態をチェックする:データ取得時(毎回)
	CheckUpdateStatus()

	// 更新状態を更新する:テーブル更新時(毎回)
	UpdateStatus()
}
