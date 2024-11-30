// api-go\src\repository\jp_daily_price.go

package repository

import "api-go/src/model"

type JpDailyPriceRepository interface {
	// 株価データを取得する:一日一回
	GetAllDailyPriceList() (*[]model.JpDailyPrice, error)

	// 株価データを追加する:一日一回
	AddDailyPriceData(*[]model.JpDailyPrice) error

	// 株価データを削除する:データが一定数に達したら(一日一回、設定した日数が経過しているレコード)
	DeleteDailyPriceData(int) error

	// 株価データを削除する:特定の日付けを受け取り、それより前の日付をもつデータを削除する
	DeleteBeforeSpecifiedDate(string) error

	// ティッカーに対応する最新の終値を取得
	GetLatestClosePricesByTickers(tickers []string) (map[string]float64, error)

	// DBに格納されている株価情報の保持する日付を返す
	GetLatestDate() (string, error)

	// その他、必要に応じて追加
}
