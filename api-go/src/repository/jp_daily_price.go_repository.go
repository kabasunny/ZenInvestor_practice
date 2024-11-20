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

	// その他、必要に応じて追加
}
