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
	tx := r.db.Begin()

	for _, price := range *newPrices {
		// 日付を日単位に統一（時刻をゼロにする）
		price.Date = time.Date(price.Date.Year(), price.Date.Month(), price.Date.Day(), 0, 0, 0, 0, price.Date.Location())

		var existingPrice model.JpDailyPrice
		err := tx.Where("ticker = ? AND date = ?", price.Ticker, price.Date).First(&existingPrice).Error
		if err == nil {
			// 重複エントリの場合スキップ
			fmt.Printf("Skipping duplicate entry for ticker: %s, date: %s\n", price.Ticker, price.Date)
			continue
		} else if err != nil && err != gorm.ErrRecordNotFound {
			tx.Rollback()
			return fmt.Errorf("failed to check existing daily price data: %w", err)
		}

		if err := tx.Create(&price).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to add daily price data: %w", err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %w", err)
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

// ティッカーに対応する最新の終値を取得
func (r *jpDailyPriceRepositoryImpl) GetLatestClosePricesByTickers(tickers []string) (map[string]float64, error) {
	var prices []model.JpDailyPrice
	if err := r.db.Where("ticker IN ?", tickers).Order("date desc").Group("ticker").Find(&prices).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch latest prices: %w", err)
	}

	priceMap := make(map[string]float64)
	for _, price := range prices {
		priceMap[price.Ticker] = price.Close
	}
	return priceMap, nil
}
