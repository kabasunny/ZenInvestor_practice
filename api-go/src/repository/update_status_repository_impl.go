// api-go\src\repository\update_status_repository_impl.go
package repository

import (
	"api-go/src/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type updateStatusRepositoryImpl struct {
	db *gorm.DB
}

func NewUpdateStatusRepository(db *gorm.DB) UpdateStatusRepository {
	return &updateStatusRepositoryImpl{db: db}
}

func (r *updateStatusRepositoryImpl) GetAllUpdateStatuses() ([]model.UpdateStatus, error) {
	var updateStatuses []model.UpdateStatus
	if err := r.db.Find(&updateStatuses).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch update statuses: %w", err)
	}
	return updateStatuses, nil
}

const (
	FieldTableName = "tb_name"
	FieldDate      = "date"
)

func (r *updateStatusRepositoryImpl) UpdateStatus(tbName string) error {
	fmt.Println("In UpdateStatus")

	// アップサート処理
	updateStatus := model.UpdateStatus{
		TbName: tbName,
		Date:   time.Now(),
	}

	result := r.db.Where(model.UpdateStatus{TbName: tbName}).
		Assign(model.UpdateStatus{Date: updateStatus.Date}).
		FirstOrCreate(&updateStatus)

	fmt.Printf("Debug - SQL Result: RowsAffected: %d, Error: %v\n", result.RowsAffected, result.Error)

	if result.Error != nil {
		return fmt.Errorf("failed to upsert update status for %s '%s': %w", FieldTableName, tbName, result.Error)
	}

	// 現在のテーブルのフィールド一覧を取得して表示（デバッグ用）
	var updateStatuses []model.UpdateStatus
	if err := r.db.Find(&updateStatuses).Error; err != nil {
		return fmt.Errorf("failed to fetch update statuses for debugging: %w", err)
	}
	for _, status := range updateStatuses {
		fmt.Printf("Debug - Table: %s, Date: %s\n", status.TbName, status.Date)
	}

	fmt.Println("Out UpdateStatus")

	return nil
}
