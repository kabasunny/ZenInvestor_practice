// api-go\src\model\update_st.go

// Update テーブルは銘柄情報および株価情報の更新状態を保存します。
// テーブル名と計算基準日を格納し、各テーブルの最新情報を管理します。

package models

import "time"

type Update struct {
	Date   time.Time `gorm:"type:date"`
	TbName string    `gorm:"primaryKey;type:text"`
}
