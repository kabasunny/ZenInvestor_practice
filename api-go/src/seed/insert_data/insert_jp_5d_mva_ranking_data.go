// api-go\src\seed\insert\insert_jp_5d_mva_ranking_data.go

package seed

import (
	"api-go/src/model"
	"time"

	"gorm.io/gorm"
)

// InsertJP5dMvaRankingData 初期データとして日本株5日間の平均売買代金ランキングを挿入
func InsertJP5dMvaRankingData(db *gorm.DB) error {
	// 現状、初期値が不要のため、テストデータの作成

	// 昨日の日付を取得
	yesterday := time.Now().AddDate(0, 0, -1)

	// テストデータの作成
	rankings := []model.Jp5dMvaRanking{
		{Ranking: 1, Ticker: "test1", Date: yesterday, AvgVolue: 100000.0},
		{Ranking: 2, Ticker: "test2", Date: yesterday, AvgVolue: 90000.0},
		{Ranking: 3, Ticker: "test3", Date: yesterday, AvgVolue: 80000.0},
		{Ranking: 4, Ticker: "test4", Date: yesterday, AvgVolue: 70000.0},
		{Ranking: 5, Ticker: "test5", Date: yesterday, AvgVolue: 60000.0},
	}

	for _, ranking := range rankings {
		if err := db.Create(&ranking).Error; err != nil {
			return err
		}
	}
	return nil
}
