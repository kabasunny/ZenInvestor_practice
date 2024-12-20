// api-go\src\repository\jp_daily_price_repository_impl.go

package repository

import (
	"api-go/src/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type jpDailyPriceRepositoryImpl struct {
	db *gorm.DB
}

func NewJpDailyPriceRepository(db *gorm.DB) JpDailyPriceRepository {
	return &jpDailyPriceRepositoryImpl{db: db}
}

// 株価データを取得する: 一日一回
func (r *jpDailyPriceRepositoryImpl) GetAllDailyPriceList() (*[]model.JpDailyPrice, error) {
	var prices []model.JpDailyPrice
	if err := r.db.Find(&prices).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch daily price list: %w", err)
	}
	return &prices, nil
}

// 株価データを追加する: 一日一回
func (r *jpDailyPriceRepositoryImpl) AddDailyPriceData(newPrices *[]model.JpDailyPrice) error {
	fmt.Println("In AddDailyPriceData")

	for _, price := range *newPrices {
		// 日付を日単位に統一（時刻をゼロにする）
		price.Date = time.Date(price.Date.Year(), price.Date.Month(), price.Date.Day(), 0, 0, 0, 0, price.Date.Location())

		// Saveメソッドはレコードが存在すれば更新し、存在しなければ挿入
		if err := r.db.Save(&price).Error; err != nil {
			return fmt.Errorf("failed to upsert daily price data for symbol: %s, date: %s: %w", price.Symbol, price.Date, err)
		}

		// fmt.Printf("Upserted price data for symbol: %s, date: %s\n", price.Symbol, price.Date)
	}

	fmt.Println("Out AddDailyPriceData")
	return nil
}

// 株価データを削除する: データが一定数に達したら（設定した日数が経過しているレコードを一日一回削除）
func (r *jpDailyPriceRepositoryImpl) DeleteDailyPriceData(days int) error {
	// 現在の日付から指定された日数を引いた日付を計算
	beforeDate := time.Now().AddDate(0, 0, -days)

	// 指定された日付以前のデータを削除
	result := r.db.Where("date < ?", beforeDate).Delete(&model.JpDailyPrice{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete daily price data: %w", result.Error)
	}

	fmt.Printf("Deleted %d records older than %s\n", result.RowsAffected, beforeDate)
	return nil
}

// 株価データを削除する: 特定の日付けを受け取り、それより前の日付をもつデータを削除する
func (r *jpDailyPriceRepositoryImpl) DeleteBeforeSpecifiedDate(specifiedDate string) error {
	fmt.Printf("In DeleteBeforeSpecifiedDate")

	// 指定された日付を time.Time 型に変換
	date, err := time.Parse("2006-01-02", specifiedDate)
	if err != nil {
		return fmt.Errorf("failed to parse specified date: %w", err)
	}

	// 指定された日付以前のデータを削除
	result := r.db.Where("date < ?", date).Delete(&model.JpDailyPrice{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete daily price data before specified date: %w", result.Error)
	}

	fmt.Printf("Deleted %d records older than %s\n", result.RowsAffected, specifiedDate)
	return nil
}

// ティッカーに対応する最新の終値を取得
func (r *jpDailyPriceRepositoryImpl) GetLatestClosePricesByTickers(symbols []string) (map[string]float64, error) {
	var results []struct {
		Symbol string
		Close  float64
	}

	if err := r.db.Table("jp_daily_price").
		Select("symbol, MAX(date) as date, MAX(close) as close").
		Where("symbol IN ?", symbols).
		Group("symbol").
		Scan(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch latest prices: %w", err)
	}

	priceMap := make(map[string]float64)
	for _, result := range results {
		priceMap[result.Symbol] = result.Close
	}
	return priceMap, nil
}

// 単純に最新の日付を取得する
func (r *jpDailyPriceRepositoryImpl) GetLatestDate() (string, error) {
	var latestPrice model.JpDailyPrice
	if err := r.db.Order("date desc").First(&latestPrice).Error; err != nil {
		return "", fmt.Errorf("failed to fetch latest date: %w", err)
	}
	return latestPrice.Date.Format("2006-01-02"), nil
}
