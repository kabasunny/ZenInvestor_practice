// api-go\src\model\rankings_5d_mva.go

// Ranking5dMva テーブルは過去5日間の平均売買代金に基づくランキング情報を保存します。
// ランキングID、銘柄コード（ティッカーシンボル）、日付を複合キーとして使用します。

package models

import "time"

type Ranking5dMva struct {
	Ranking  int       `gorm:"primaryKey"`
	Ticker   string    `gorm:"primaryKey;type:text"`
	Date     time.Time `gorm:"primaryKey;type:date"`
	AvgVolue float64   `gorm:"type:real"`
}
