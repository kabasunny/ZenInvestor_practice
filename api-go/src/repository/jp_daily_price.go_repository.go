// api-go\src\repository\jp_daily_price.go_repository.go

package repository

type JpDailyPriceRepository interface {
	// 株価データを取得する:一日一回
	GetPriceData()

	// 株価データを追加する:一日一回
	AddPriceData()

	// 株価データを削除する:データが一定数に達したら(一日一回)
	DeletePriceData()
}
