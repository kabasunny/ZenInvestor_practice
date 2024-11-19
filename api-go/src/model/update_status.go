// api-go\src\model\update_status.go

// Update テーブルは銘柄情報および株価情報の更新状態を保存。 011_create_update_table.go
// テーブル名と計算基準日を格納し、各テーブルの最新情報を管理。

package model

import "time"

type UpdateStatus struct {
	Date   time.Time `gorm:"type:date"`
	TbName string    `gorm:"primaryKey;type:varchar(20)"`
}

// TableName カスタムテーブル名を指定するメソッド
// GORMを使用してGoでMySQLにテーブルを作成すると、デフォルトでテーブル名は構造体名のスネークケース形式になり、複数形になる
func (UpdateStatus) TableName() string {
	return "update_status"
}
