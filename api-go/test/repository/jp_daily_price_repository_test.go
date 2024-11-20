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
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	// 新しいデータの挿入
	newPrices := []model.JpDailyPrice{
		{Ticker: "test_add", Date: time.Now().Truncate(24 * time.Hour), Open: 1000.0, Close: 1100.0, High: 1150.0, Low: 950.0, Volume: 10000, Volue: 1000000},
	}

	err := repo.AddDailyPriceData(&newPrices)
	assert.NoError(t, err)

	// 追加後のデータを取得して確認
	var addedPrices []model.JpDailyPrice
	db.Find(&addedPrices)
	fmt.Println("Added DailyPriceData:")
	for _, price := range addedPrices {
		fmt.Println(price)
	}

	// 追加したデータの削除 (日付を指定) - clause.Returning{} を追加
	deleteDate := newPrices[0].Date
	deleteResult := db.Clauses(clause.Returning{}).Delete(&model.JpDailyPrice{}, "date = ? AND ticker = ?", deleteDate, newPrices[0].Ticker)
	if deleteResult.Error != nil {
		t.Fatal("failed to delete record:", deleteResult.Error)
	}

	// 削除後のデータを再取得して確認 - db.Where を使用してクエリ
	var afterDelete []model.JpDailyPrice
	db.Where("ticker = ? AND date = ?", newPrices[0].Ticker, deleteDate).Find(&afterDelete)

	// 削除が行われたことを確認 - RecordNotFound エラーを期待
	var deletedResult model.JpDailyPrice
	err = db.First(&deletedResult, "ticker = ? AND date = ?", newPrices[0].Ticker, deleteDate).Error
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
}

func TestDeleteDailyPriceData(t *testing.T) {
	db := repository_test_helper.SetupTestDB()

	repo := repository.NewJpDailyPriceRepository(db)

	// 古い日付のデータを作成
	oldDate := time.Now().AddDate(0, 0, -31) // 31日前の日付
	newPrices := []model.JpDailyPrice{
		{Ticker: "test_delete", Date: oldDate, Open: 1000.0, Close: 1100.0, High: 1150.0, Low: 950.0, Volume: 10000, Volue: 1000000},
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

	// 30日以上前の日付のデータを削除
	err = repo.DeleteDailyPriceData(30)
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
