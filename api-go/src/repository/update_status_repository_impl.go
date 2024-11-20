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
	result := r.db.Model(&model.UpdateStatus{}).
		Where(fmt.Sprintf("%s = ?", FieldTableName), tbName).
		Update(FieldDate, time.Now())

	if result.Error != nil {
		return fmt.Errorf("failed to update status for %s '%s': %w", FieldTableName, tbName, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no records found with %s '%s'", FieldTableName, tbName)
	}

	return nil
}
