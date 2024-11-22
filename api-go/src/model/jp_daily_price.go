// api-go\src\model\jp_daily_price.go

// DailyPrice テーブルは各銘柄の日次株価データを保存。002_create_daily_prices_table.go
// 銘柄コード（ティッカーシンボル）と日付を複合キーとして使用。

package model

import "time"

type JpDailyPrice struct {
	Ticker   string    `gorm:"primaryKey;type:varchar(10)"` // VARCHAR型に変更し、長さを指定
	Date     time.Time `gorm:"primaryKey;type:date"`
	Open     float64   `gorm:"type:real"`
	Close    float64   `gorm:"type:real"`
	High     float64   `gorm:"type:real"`
	Low      float64   `gorm:"type:real"`
	Volume   int64     `gorm:"type:integer"`
	Turnover int64     `gorm:"type:integer"`
}

// TableName カスタムテーブル名を指定するメソッド
// GORMを使用してGoでMySQLにテーブルを作成すると、デフォルトでテーブル名は構造体名のスネークケース形式になり、複数形になる
func (JpDailyPrice) TableName() string {
	return "jp_daily_price"
}
