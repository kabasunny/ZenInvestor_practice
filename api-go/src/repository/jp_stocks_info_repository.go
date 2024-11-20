// api-go\src\repository\jp_stocks_info.go

package repository

import "api-go/src/model"

type JpStockInfoRepository interface {
	// 銘柄データを取得する:一日一回
	GetAllStockInfo() (*[]model.JpStockInfo, error)

	// GetAllStockInfoByTick(tickerCode string) (*model.JpStockInfo, error)

	// 銘柄データを更新する:一日一回
	UpdateStockInfo(newJpStockInfo *[]model.JpStockInfo) error

	// UpdateStockInfo(newJpStockInfo *[]model.JpStockInfo) error
	// 銘柄データをすべて削除する:使用未定
	// DeleteAllStockInfo() error

	// 銘柄を指定して削除する:使用未定
	// DeleteStockInfoByTick(tickerCode string) error
}
