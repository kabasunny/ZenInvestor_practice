// api-go\src\repository\update_status_repository_impl.go

package repository

import (
	"api-go/src/model"

	"gorm.io/gorm"
)

type updateStatusRepositoryImpl struct {
	db *gorm.DB
}

func NewUpdateStatusRepository(db *gorm.DB) UpdateStatusRepository {
	return &updateStatusRepositoryImpl{db: db}
}

func (repo *updateStatusRepositoryImpl) GetAllUpdateStatuses() ([]model.UpdateStatus, error) {
	var updateStatuses []model.UpdateStatus
	if err := repo.db.Find(&updateStatuses).Error; err != nil {
		return nil, err
	}
	return updateStatuses, nil
}

func (repo *updateStatusRepositoryImpl) UpdateStatus(updateStatus *model.UpdateStatus) error {
	return repo.db.Save(updateStatus).Error
}
