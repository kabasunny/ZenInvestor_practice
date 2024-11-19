// api-go\src\seed\seed_data.go

package seed

import (
	"gorm.io/gorm"
)

// InsertInitialData 各テーブルの初期データを挿入する関数をまとめて呼び出します
func InsertInitialData(db *gorm.DB) error {
	if err := InsertJPStocksInfoData(db); err != nil {
		return err
	}
	if err := InsertJPDailyPricesData(db); err != nil {
		return err
	}
	if err := InsertUpdateStatusData(db); err != nil {
		return err
	}
	if err := InsertJP5dMvaRankingData(db); err != nil {
		return err
	}
	return nil
}
