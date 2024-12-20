// api-go\src\seed\insert\insert_jp_daily_price_data.go
package seed

import (
	"api-go/src/model"
	"time"

	"gorm.io/gorm"
)

// InsertJPDailyPricesData 初期データとして日本株日次価格を挿入
func InsertJPDailyPricesData(db *gorm.DB) error {
	// 現状、初期値が不要のため、テストデータの作成

	// 昨日の日付を取得
	yesterday := time.Now().AddDate(0, 0, -1)

	// テストデータの作成
	dailyPrices := []model.JpDailyPrice{
		{Symbol: "test1", Date: yesterday, Open: 100.0, Close: 110.0, High: 115.0, Low: 95.0, Volume: 10000, Turnover: 1000000},
		{Symbol: "test2", Date: yesterday, Open: 200.0, Close: 210.0, High: 215.0, Low: 195.0, Volume: 20000, Turnover: 2000000},
		{Symbol: "test3", Date: yesterday, Open: 300.0, Close: 310.0, High: 315.0, Low: 295.0, Volume: 30000, Turnover: 3000000},
		{Symbol: "test4", Date: yesterday, Open: 400.0, Close: 410.0, High: 415.0, Low: 395.0, Volume: 40000, Turnover: 4000000},
		{Symbol: "test5", Date: yesterday, Open: 500.0, Close: 510.0, High: 515.0, Low: 495.0, Volume: 50000, Turnover: 5000000},
	}

	for _, price := range dailyPrices {
		if err := db.Create(&price).Error; err != nil {
			return err
		}
	}
	return nil
}
