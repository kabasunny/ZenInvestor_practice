// api-go\src\model\jp_5d_mva_ranking.go

// Ranking5dMva テーブルは過去5日間の平均売買代金に基づくランキング情報を保存。 021_create_srankings_5d_mva_table.go
// ランキングID、銘柄コード（ティッカーシンボル）、日付を複合キーとして使用。

package model

import "time"

type Jp5dMvaRanking struct {
	Ranking  int       `gorm:"primaryKey"`
	Ticker   string    `gorm:"primaryKey;type:varchar(10)"`
	Date     time.Time `gorm:"primaryKey;type:date"`
	AvgVolue float64   `gorm:"type:real"`
}

// TableName カスタムテーブル名を指定するメソッド
// GORMを使用してGoでMySQLにテーブルを作成すると、デフォルトでテーブル名は構造体名のスネークケース形式になり、複数形になる
func (Jp5dMvaRanking) TableName() string {
	return "jp_5d_mva_ranking"
}
