// api-go\src\model\daily_prices.go

// DailyPrice テーブルは各銘柄の日次株価データを保存します。
// 銘柄コード（ティッカーシンボル）と日付を複合キーとして使用します。

package models

import "time"

type DailyPrice struct {
	Ticker string    `gorm:"primaryKey;type:text"`
	Date   time.Time `gorm:"primaryKey;type:date"`
	Open   float64   `gorm:"type:real"`
	Close  float64   `gorm:"type:real"`
	High   float64   `gorm:"type:real"`
	Low    float64   `gorm:"type:real"`
	Volume int64     `gorm:"type:integer"`
	Volue  int64     `gorm:"type:integer"`
}
