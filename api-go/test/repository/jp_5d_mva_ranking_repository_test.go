// api-go\test\repository\jp_5d_mva_ranking_repository_test.go
package repository_test

import (
	"api-go/src/model"
	"api-go/src/repository"
	"api-go/test/repository/repository_test_helper"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGet5dMvaRankingData(t *testing.T) {
	db := repository_test_helper.SetupTestDB()
	repo := repository.NewJp5dMvaRankingRepository(db)

	rankings, err := repo.Get5dMvaRankingData()
	assert.NoError(t, err)

	fmt.Println("Retrieved 5d MVA Ranking Data:")
	for _, ranking := range *rankings {
		fmt.Println(ranking)
	}
}

func TestAdd5dMvaRankingData(t *testing.T) {
	db := repository_test_helper.SetupTestDB()
	repo := repository.NewJp5dMvaRankingRepository(db)

	// 既存のデータを削除
	db.Exec("DELETE FROM jp_daily_price WHERE ticker IN ('test6', 'test7')")

	// 本日から7日分の2銘柄分のダミーデータを作成し、jp_daily_priceテーブルに挿入する
	today := time.Now().Truncate(24 * time.Hour)
	ticker1 := "test6"
	ticker2 := "test7"
	dummyPrices := []model.JpDailyPrice{
		{Ticker: ticker1, Date: today.AddDate(0, 0, -6), Turnover: 1000},
		{Ticker: ticker1, Date: today.AddDate(0, 0, -5), Turnover: 1100},
		{Ticker: ticker1, Date: today.AddDate(0, 0, -4), Turnover: 1200},
		{Ticker: ticker1, Date: today.AddDate(0, 0, -3), Turnover: 1300},
		{Ticker: ticker1, Date: today.AddDate(0, 0, -2), Turnover: 1400},
		{Ticker: ticker1, Date: today.AddDate(0, 0, -1), Turnover: 1500},
		{Ticker: ticker2, Date: today.AddDate(0, 0, -6), Turnover: 2000},
		{Ticker: ticker2, Date: today.AddDate(0, 0, -5), Turnover: 2100},
		{Ticker: ticker2, Date: today.AddDate(0, 0, -4), Turnover: 2200},
		{Ticker: ticker2, Date: today.AddDate(0, 0, -3), Turnover: 2300},
		{Ticker: ticker2, Date: today.AddDate(0, 0, -2), Turnover: 2400},
		{Ticker: ticker2, Date: today.AddDate(0, 0, -1), Turnover: 2500},
	}

	for _, price := range dummyPrices {
		err := db.Create(&price).Error
		assert.NoError(t, err)
	}

	// 挿入後のjp_daily_priceテーブルの内容を表示する
	var dailyPrices []model.JpDailyPrice
	db.Find(&dailyPrices)
	fmt.Println("jp_daily_price Data:")
	for _, price := range dailyPrices {
		fmt.Println(price)
	}

	// Add5dMvaRankingDataを呼び出す
	err := repo.Add5dMvaRankingData()
	assert.NoError(t, err)

	// ランキングデータが挿入されていることを確認するために、jp_5d_mva_rankingテーブルの内容を表示する
	var rankings []model.Jp5dMvaRanking
	db.Find(&rankings)
	fmt.Println("jp_5d_mva_ranking Data:")
	for _, ranking := range rankings {
		fmt.Println(ranking)
	}

	// jp_daily_priceテーブルに挿入したデータを削除する
	db.Exec("DELETE FROM jp_daily_price WHERE ticker IN ('test6', 'test7')")

	// jp_5d_mva_rankingテーブルに挿入したデータを本日付で削除する リアル日付の経過とともにjp_daily_priceにある古いデータが無尽蔵に増える
	db.Exec("DELETE FROM jp_5d_mva_ranking WHERE date = ?", today)
}

func TestDelete5dMvaRankingData(t *testing.T) {
	db := repository_test_helper.SetupTestDB()
	repo := repository.NewJp5dMvaRankingRepository(db)

	// 古い日付のデータを作成
	oldDate := time.Now().AddDate(0, 0, -31).Truncate(24 * time.Hour)
	newRankings := []model.Jp5dMvaRanking{
		{Ranking: 1, Ticker: "test_tic", Date: oldDate, AvgTurnover: 12345.67},
	}

	// ダミーデータを jp_5d_mva_ranking テーブルに直接挿入
	for _, ranking := range newRankings {
		err := db.Create(&ranking).Error
		assert.NoError(t, err)
	}

	// 削除前のデータ確認
	var beforeDelete []model.Jp5dMvaRanking
	db.Find(&beforeDelete)
	fmt.Println("Before Delete 5d MVA Ranking Data:")
	for _, ranking := range beforeDelete {
		fmt.Println(ranking)
	}

	// 10日以上前の日付のデータを削除
	err := repo.Delete5dMvaRankingData(10)
	assert.NoError(t, err)

	// 削除後のデータ確認
	var afterDelete []model.Jp5dMvaRanking
	db.Find(&afterDelete)
	fmt.Println("After Delete 5d MVA Ranking Data:")
	for _, ranking := range afterDelete {
		fmt.Println(ranking)
	}

	// 確認：削除されたデータが存在しないこと
	var deletedResult model.Jp5dMvaRanking
	err = db.First(&deletedResult, "ranking = ? AND ticker = ? AND date = ?", 1, "test_ticker", oldDate).Error
	assert.Error(t, err) // レコードが見つからないエラーを期待
}

// go test -v ./test/repository/jp_5d_mva_ranking_repository_test.go -run TestGet5dMvaRankingData
// go test -v ./test/repository/jp_5d_mva_ranking_repository_test.go -run TestAdd5dMvaRankingData
// go test -v ./test/repository/jp_5d_mva_ranking_repository_test.go -run TestDelete5dMvaRankingData
