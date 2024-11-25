// api-go\src\repository\jp_stocks_info_repository_impl.go
package repository

import (
	"api-go/src/model"
	"fmt"

	"gorm.io/gorm"
)

type jpStockInfoRepositoryImpl struct {
	db *gorm.DB
}

func NewJpStockInfoRepository(db *gorm.DB) JpStockInfoRepository {
	return &jpStockInfoRepositoryImpl{db: db}
}

// 銘柄データを取得する: 一日一回
func (r *jpStockInfoRepositoryImpl) GetAllStockInfo() (*[]model.JpStockInfo, error) {
	var stocks []model.JpStockInfo
	if err := r.db.Find(&stocks).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch stock info: %w", err)
	}
	return &stocks, nil
}

// 銘柄データを更新する: 一日一回
func (r *jpStockInfoRepositoryImpl) UpdateStockInfo(newJpStockInfo *[]model.JpStockInfo) error {
	fmt.Println("In UpdateStockInfo")

	for _, stock := range *newJpStockInfo {
		// Saveメソッドを使用してアップサート処理を実行
		if err := r.db.Save(&stock).Error; err != nil {
			return fmt.Errorf("failed to upsert stock info for ticker '%s': %w", stock.Ticker, err)
		}

		fmt.Printf("Upserted stock info for ticker: %s\n", stock.Ticker)
	}

	fmt.Println("Out UpdateStockInfo")
	return nil
}

// 銘柄データをすべて削除する
func (r *jpStockInfoRepositoryImpl) DeleteAllStockInfo() error {
	if err := r.db.Exec("DELETE FROM jp_stocks_info").Error; err != nil {
		return fmt.Errorf("failed to delete all stock info: %w", err)
	}
	return nil
}

// 新しい銘柄データを挿入する
func (r *jpStockInfoRepositoryImpl) InsertStockInfo(newJpStockInfo *[]model.JpStockInfo) error {
	tx := r.db.Begin()
	if err := tx.Create(newJpStockInfo).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert stock info: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

// ティッカーに対応する銘柄情報を取得
func (r *jpStockInfoRepositoryImpl) GetStockInfoByTickers(tickers []string) (map[string]model.JpStockInfo, error) {
	var stockList []model.JpStockInfo
	if err := r.db.Where("ticker IN ?", tickers).Find(&stockList).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch stock info: %w", err)
	}

	stocks := make(map[string]model.JpStockInfo)
	for _, stock := range stockList {
		stocks[stock.Ticker] = stock
	}
	return stocks, nil
}
