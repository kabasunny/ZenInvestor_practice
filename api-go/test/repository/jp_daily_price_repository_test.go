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
		{Ticker: ticker, Date: date, Open: 1000.0, Close: 1100.0, High: 1150.0, Low: 950.0, Volume: 10000, Value: 1000000},
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
	db.Where("ticker = ? AND DATE(date) = ?", addPrices[0].Ticker, dateString).Delete(&model.JpDailyPrice{})

}

func TestDeleteDailyPriceData(t *testing.T) {
	db := repository_test_helper.SetupTestDB()

	repo := repository.NewJpDailyPriceRepository(db)

	// 古い日付のデータを作成
	oldDate := time.Now().AddDate(0, 0, -31) // 31日前の日付
	newPrices := []model.JpDailyPrice{
		{Ticker: "testdel", Date: oldDate, Open: 1000.0, Close: 1100.0, High: 1150.0, Low: 950.0, Volume: 10000, Value: 1000000},
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
	err = db.First(&deletedResult, "ticker = ? AND date = ?", "test_delete", oldDate).Error
	assert.Error(t, err) // レコードが見つからないエラーを期待
}

// テストの実行コード
// go test -v ./test/repository/jp_daily_price_repository_test.go -run TestGetAllDailyPriceList
// go test -v ./test/repository/jp_daily_price_repository_test.go -run TestAddDailyPriceData
// go test -v ./test/repository/jp_daily_price_repository_test.go -run TestDeleteDailyPriceData
