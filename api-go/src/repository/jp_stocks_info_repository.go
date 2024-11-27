// api-go\src\repository\jp_stocks_info.go

package repository

import "api-go/src/model"

type JpStockInfoRepository interface {
	// 銘柄データ一覧を取得する:一日一回
	GetAllStockInfo() (*[]model.JpStockInfo, error)

	// 銘柄コード一覧を取得する
	GetAllSymbols() ([]string, error)

	// 銘柄データを更新する:一日一回
	UpdateStockInfo(newJpStockInfo *[]model.JpStockInfo) error

	// ティッカーに対応する銘柄情報を取得
	GetStockInfoByTickers(tickers []string) (map[string]model.JpStockInfo, error)

	// UpdateStockInfo(newJpStockInfo *[]model.JpStockInfo) error

	// 銘柄データをすべて削除する
	DeleteAllStockInfo() error

	// 銘柄データを更新する:一日一回
	InsertStockInfo(newJpStockInfo *[]model.JpStockInfo) error

	// 銘柄を指定して削除する:使用未定
	// DeleteStockInfoByTick(tickerCode string) error
}
