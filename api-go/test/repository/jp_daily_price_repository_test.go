// api-go\test\repository\jp_daily_price_repository_test.go
package repository

import (
	"api-go/src/model"
	"api-go/src/repository"
	"api-go/test/repository/repository_test_helper"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	// "gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

func TestGetAllDailyPriceList(t *testing.T) {
	db := repository_test_helper.SetupTestDB()

	repo := repository.NewJpDailyPriceRepository(db)

	// すべての日次株価データを取得する
	prices, err := repo.GetAllDailyPriceList()
	assert.NoError(t, err)
	assert.NotNil(t, prices)

	// 結果をターミナルに表示
	fmt.Println("GetAllDailyPriceList:")
	for _, price := range *prices {
		fmt.Println(price)
	}
}

func TestAddDailyPriceData(t *testing.T) {
	db := repository_test_helper.SetupTestDB()

	repo := repository.NewJpDailyPriceRepository(db)

	date := time.Now().Truncate(24 * time.Hour)
	fmt.Println("test_add Date:", date)
	ticker := "test_add"
	fmt.Println("test_add Ticker:", ticker)
	// 新しいデータの挿入
	addPrices := []model.JpDailyPrice{
		{Symbol: ticker, Date: date, Open: 1000.0, Close: 1100.0, High: 1150.0, Low: 950.0, Volume: 10000, Turnover: 1000000},
	}

	err := repo.AddDailyPriceData(&addPrices)
	assert.NoError(t, err)

	// 追加後のデータを取得して確認
	var addedPrices []model.JpDailyPrice
	db.Find(&addedPrices)
	fmt.Println("Added DailyPriceData:")
	for _, price := range addedPrices {
		fmt.Println(price)
	}

	// 追加したデータの削除 (日付部分のみで判定)
	dateString := addPrices[0].Date.Format("2006-01-02")
	db.Where("symbol = ? AND DATE(date) = ?", addPrices[0].Symbol, dateString).Delete(&model.JpDailyPrice{})
}

func TestDeleteDailyPriceData(t *testing.T) {
	db := repository_test_helper.SetupTestDB()

	repo := repository.NewJpDailyPriceRepository(db)

	// 古い日付のデータを作成
	oldDate := time.Now().AddDate(0, 0, -31) // 31日前の日付
	newPrices := []model.JpDailyPrice{
		{Symbol: "testdel", Date: oldDate, Open: 1000.0, Close: 1100.0, High: 1150.0, Low: 950.0, Volume: 10000, Turnover: 1000000},
	}

	err := repo.AddDailyPriceData(&newPrices)
	assert.NoError(t, err)

	// 削除前のデータ確認
	var beforeDelete []model.JpDailyPrice
	db.Find(&beforeDelete)
	fmt.Println("Before Delete DailyPriceData:")
	for _, price := range beforeDelete {
		fmt.Println(price)
	}

	// 10日以上前の日付のデータを削除
	err = repo.DeleteDailyPriceData(10)
	assert.NoError(t, err)

	// 削除後のデータ確認
	var afterDelete []model.JpDailyPrice
	db.Find(&afterDelete)
	fmt.Println("After Delete DailyPriceData:")
	for _, price := range afterDelete {
		fmt.Println(price)
	}

	// 確認：削除されたデータが存在しないこと
	var deletedResult model.JpDailyPrice
	err = db.First(&deletedResult, "symbol = ? AND date = ?", "test_delete", oldDate).Error
	assert.Error(t, err) // レコードが見つからないエラーを期待
}

func TestGetLatestClosePricesByTickers(t *testing.T) {
	db := repository_test_helper.SetupTestDB()

	repo := repository.NewJpDailyPriceRepository(db)

	// テストデータの追加
	date1 := time.Now().AddDate(0, 0, -2).Truncate(24 * time.Hour) // 2日前の日付
	date2 := time.Now().AddDate(0, 0, -1).Truncate(24 * time.Hour) // 1日前の日付

	testPrices := []model.JpDailyPrice{
		{Symbol: "test_21", Date: date1, Close: 100.0},
		{Symbol: "test_21", Date: date2, Close: 110.0},
		{Symbol: "test_22", Date: date1, Close: 200.0},
		{Symbol: "test_22", Date: date2, Close: 210.0},
	}

	err := repo.AddDailyPriceData(&testPrices)
	assert.NoError(t, err)

	// 最新の終値を取得する
	tickers := []string{"test_21", "test_22"}
	prices, err := repo.GetLatestClosePricesByTickers(tickers)
	assert.NoError(t, err)

	// 結果をターミナルに表示
	fmt.Println("GetLatestClosePricesByTickers:")
	for ticker, price := range prices {
		fmt.Printf("Ticker: %s, Latest Close Price: %.2f\n", ticker, price)
	}

	// 結果の確認
	assert.Equal(t, 2, len(prices))
	assert.Equal(t, 110.0, prices["test_21"])
	assert.Equal(t, 210.0, prices["test_22"])

	// クリーンアップ: 追加したデータを削除
	dateStrings := []string{date1.Format("2006-01-02"), date2.Format("2006-01-02")}
	db.Where("ticker IN ? AND DATE(date) IN ?", tickers, dateStrings).Delete(&model.JpDailyPrice{})
}

// テストの実行コード
// go test -v ./test/repository/jp_daily_price_repository_test.go -run TestGetAllDailyPriceList
// go test -v ./test/repository/jp_daily_price_repository_test.go -run TestAddDailyPriceData
// go test -v ./test/repository/jp_daily_price_repository_test.go -run TestDeleteDailyPriceData
// go test -v ./test/repository/jp_daily_price_repository_test.go -run TestGetLatestClosePricesByTickers
