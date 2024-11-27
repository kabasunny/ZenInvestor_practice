// api-go\test\repository\jp_stock_info_repository_test.go
package repository

import (
	"api-go/src/model"
	"api-go/src/repository"
	"api-go/test/repository/repository_test_helper"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllStockInfo(t *testing.T) {
	db := repository_test_helper.SetupTestDB()

	repo := repository.NewJpStockInfoRepository(db)

	// すべてのストック情報を取得する
	stocks, err := repo.GetAllStockInfo()
	assert.NoError(t, err)
	assert.NotNil(t, stocks)

	// 結果をターミナルに表示
	fmt.Println("GetAllStockInfo:")
	for _, stock := range *stocks {
		fmt.Println(stock)
	}
}

func TestGetAllSymbols(t *testing.T) {
	db := repository_test_helper.SetupTestDB()

	repo := repository.NewJpStockInfoRepository(db)

	// 新しいデータの挿入
	addStocks := []model.JpStockInfo{
		{Symbol: "ticker1", Name: "Stock 1", Sector: "Sector 1", Industry: "Industry 1"},
		{Symbol: "ticker2", Name: "Stock 2", Sector: "Sector 2", Industry: "Industry 2"},
	}
	db.Create(&addStocks)

	// すべてのシンボルを取得する
	symbols, err := repo.GetAllSymbols()
	assert.NoError(t, err)
	assert.NotNil(t, symbols)

	// 結果をターミナルに表示
	fmt.Println("GetAllSymbols:")
	for _, symbol := range symbols {
		fmt.Println(symbol)
	}

	// シンボル数の確認
	// assert.Equal(t, len(addStocks), len(symbols))

	// 追加したデータの削除
	db.Where("symbol IN ?", []string{"ticker1", "ticker2"}).Delete(&model.JpStockInfo{})
}

func TestUpdateStockInfo(t *testing.T) {
	db := repository_test_helper.SetupTestDB()

	repo := repository.NewJpStockInfoRepository(db)

	// 新しいデータの挿入
	addStock := model.JpStockInfo{Symbol: "test_add", Name: "Bf Stock", Sector: "Bf Sector", Industry: "Bf Industry"}
	db.Create(&addStock)

	// 追加後のデータを取得して表示
	var beforeUpdate []model.JpStockInfo
	db.Find(&beforeUpdate)
	fmt.Println("Before Update:")
	for _, stock := range beforeUpdate {
		fmt.Println(stock)
	}

	// 追加したストック情報の更新
	updatedStockInfo := []model.JpStockInfo{
		{Symbol: "test_add", Name: "Af Stock", Sector: "Af Sector", Industry: "Af Industry"},
	}
	err := repo.UpdateStockInfo(&updatedStockInfo)
	assert.NoError(t, err)

	// 更新後のデータを取得して表示
	var afterUpdate []model.JpStockInfo
	db.Find(&afterUpdate)
	fmt.Println("After Update:")
	for _, stock := range afterUpdate {
		fmt.Println(stock)
	}

	// 更新された新しいデータが正しく更新されたことを確認
	var updatedResult model.JpStockInfo
	db.First(&updatedResult, "symbol = ?", "test_add")
	assert.Equal(t, "Af Stock", updatedResult.Name)
	assert.Equal(t, "Af Sector", updatedResult.Sector)
	assert.Equal(t, "Af Industry", updatedResult.Industry)

	// 追加したデータの削除
	db.Delete(&addStock)
}

// テストの実行コード
// go test -v ./test/repository/jp_stock_info_repository_test.go -run TestGetAllStockInfo
// go test -v ./test/repository/jp_stock_info_repository_test.go -run TestGetAllSymbols
// go test -v ./test/repository/jp_stock_info_repository_test.go -run TestUpdateStockInfo
