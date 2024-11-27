// api-go\src\seed\insert\insert_jp_stocks_info_data.go
package seed

import (
	"api-go/src/model"
	"log"

	"gorm.io/gorm"
)

// InsertJPStocksInfoData 初期データとして日本株情報を挿入
func InsertJPStocksInfoData(db *gorm.DB) error {
	// 現状、初期値が不要のため、テストデータの作成
	stocksInfo := []model.JpStockInfo{
		{Symbol: "test1", Name: "Test Stock 1", Sector: "Technology", Industry: "Software"},
		{Symbol: "test2", Name: "Test Stock 2", Sector: "Healthcare", Industry: "Pharmaceuticals"},
		{Symbol: "test3", Name: "Test Stock 3", Sector: "Finance", Industry: "Banking"},
		{Symbol: "test4", Name: "Test Stock 4", Sector: "Energy", Industry: "Oil & Gas"},
		{Symbol: "test5", Name: "Test Stock 5", Sector: "Consumer Goods", Industry: "Retail"},
	}

	for _, stock := range stocksInfo {
		var existing model.JpStockInfo
		if err := db.Where("Symbol = ?", stock.Symbol).First(&existing).Error; err == nil {
			// エントリが既に存在する場合はスキップ
			log.Printf("Skipping insert for existing Symbol: %v", stock.Symbol)
			continue
		}

		// エントリが存在しない場合のみ挿入
		if err := db.Create(&stock).Error; err != nil {
			return err
		}
	}
	return nil
}
