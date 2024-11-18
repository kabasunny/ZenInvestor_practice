// api-go\src\model\stocks_info.go

// StockInfo テーブルは銘柄の基本情報を保存します。
// 銘柄コード（ティッカーシンボル）を主キーとし、他のテーブルとリレーションを構築します。

package models

import "time"

type StockInfo struct {
	Ticker      string    `gorm:"primaryKey;type:text"`
	Name        string    `gorm:"type:text"`
	Sector      string    `gorm:"type:text"`
	Industry    string    `gorm:"type:text"`
	MarketCap   float64   `gorm:"type:real"`
	ListingDate time.Time `gorm:"type:date"`
}
