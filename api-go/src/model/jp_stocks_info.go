// api-go\src\model\jp_stocks_info.go

package model

// StockInfo テーブルは銘柄の基本情報を保存。001_create_stocks_info_table.go
// 銘柄コード（ティッカーシンボル）を主キーとし、他のテーブルとリレーションを構築。

type JpStockInfo struct {
	Symbol   string `gorm:"primaryKey;type:varchar(10)"` // VARCHAR型に変更し、長さを10に設定
	Name     string `gorm:"type:text"`
	Sector   string `gorm:"type:text"`
	Industry string `gorm:"type:text"`
	Date     string `gorm:"type:text"`
	// MarketCap   float64   `gorm:"type:real"` // 無しにする
	// ListingDate time.Time `gorm:"type:date"` // 無しにする
}

// TableName カスタムテーブル名を指定するメソッド
// GORMを使用してGoでMySQLにテーブルを作成すると、デフォルトでテーブル名は構造体名のスネークケース形式になり、複数形になる
func (JpStockInfo) TableName() string {
	return "jp_stocks_info"
}
