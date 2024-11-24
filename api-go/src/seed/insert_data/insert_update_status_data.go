// api-go\src\seed\insert\insert_update_status_data.go

package seed

import (
	"api-go/src/model"
	"time"

	"gorm.io/gorm"
)

func InsertUpdateStatusData(db *gorm.DB) error {
	yesterday := time.Now().AddDate(0, 0, -1)

	// 各テーブルの更新状態を挿入
	updateStatuses := []model.UpdateStatus{
		// 初期化は昨日の日付で行う
		// 不要　{Date: yesterday, TbName: "update_status"},
		{Date: yesterday, TbName: "jp_stocks_info"},
		{Date: yesterday, TbName: "jp_daily_price"},
		{Date: yesterday, TbName: "jp_5d_mva_ranking"},
	}

	for _, status := range updateStatuses {
		if err := db.Create(&status).Error; err != nil {
			return err
		}
	}
	return nil
}
