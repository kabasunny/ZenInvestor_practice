// api-go\src\model\jp_trading_calendar.go.go

// DailyPrice テーブルは各銘柄の日次株価データを保存。002_create_daily_prices_table.go
// 銘柄コード（ティッカーシンボル）と日付を複合キーとして使用。

package model

type JpTradingCalender struct {
	Date     string `gorm:"primaryKey;type:varchar(10)"` // 文字列でやり取りするのがよさそうなので、文字列で
	Division string `gorm:"type:char"`                   // とりあえず、ここも文字で　非営業日:0	営業日:1 東証半日立会日:2 非営業日(祝日取引あり):3
}

// TableName カスタムテーブル名を指定するメソッド
// GORMを使用してGoでMySQLにテーブルを作成すると、デフォルトでテーブル名は構造体名のスネークケース形式になり、複数形になる
func (JpTradingCalender) TableName() string {
	return "jp_trading_calendar"
}
