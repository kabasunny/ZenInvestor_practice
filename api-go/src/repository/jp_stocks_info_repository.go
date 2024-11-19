// api-go\src\repository\jp_stocks_info.go

package repository

type JpStockSInfoRepository interface {
	// 銘柄データを取得する:一日一回
	GetStockData()

	// 銘柄データを更新する:一日一回
	UpdateStockData()

	// 銘柄データを更新する:使用未定
	OtherUpdateStockData()
}
