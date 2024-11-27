// api-go\src\service\ranking_service_impl.go

package service

import (
	"api-go/src/dto"
	"api-go/src/model"
	"api-go/src/repository"
	"context"
	"fmt"
	"time"
)

// RankingServiceImpl は RankingService インターフェースの実装
type RankingServiceImpl struct {
	udsRepo  repository.UpdateStatusRepository
	jsiRepo  repository.JpStockInfoRepository
	jdpRepo  repository.JpDailyPriceRepository
	j5mrRepo repository.JP5dMvaRankingRepository
	clients  map[string]interface{}
}

// NewRankingService は新しい RankingService インスタンスを作成
func NewRankingService(
	udsRepo repository.UpdateStatusRepository,
	jsiRepo repository.JpStockInfoRepository,
	jdpRepo repository.JpDailyPriceRepository,
	j5mrRepo repository.JP5dMvaRankingRepository,
	clients map[string]interface{},
) RankingService {
	return &RankingServiceImpl{
		udsRepo:  udsRepo,
		jsiRepo:  jsiRepo,
		jdpRepo:  jdpRepo,
		j5mrRepo: j5mrRepo,
		clients:  clients,
	}
}

// GetRankingData はランキングデータを取得し、DTO にマッピング
func (s *RankingServiceImpl) GetRankingData(ctx context.Context) (*[]dto.RankingServiceResponse, error) {
	fmt.Println("In GetRankingData")
	// update_statusテーブルの構造体を取得
	statuses, err := s.udsRepo.GetAllUpdateStatuses()
	fmt.Println(statuses)
	if err != nil {
		return nil, fmt.Errorf("failed to get update statuses: %w", err)
	}

	// 今日の日付を取得
	today := time.Now()
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	fmt.Println("Got Today : today =", today)

	// 各テーブルの更新日を格納する変数を初期化
	jp5dMvaRankingDate, _, _ := getLatestUpdateDates(statuses)
	fmt.Println("jp5dMvaRankingDate =", jp5dMvaRankingDate)

	// 状態に応じたデータの取得
	if jp5dMvaRankingDate.Equal(today) {
		fmt.Println("全てのデータがそろっているため、ランキング取得")
		return s.fetchRankingData()
	} else {
		fmt.Println("データが最新ではないため、バッチ処理が必要")
		return nil, fmt.Errorf("data is not up-to-date. Please run batch processes.")
	}
}

// getLatestUpdateDates は各テーブルの更新日を取得
func getLatestUpdateDates(statuses []model.UpdateStatus) (time.Time, time.Time, time.Time) {
	fmt.Println("In getLatestUpdateDates")

	var jp5dMvaRankingDate, jpDailyPriceDate, jpStocksInfoDate time.Time
	for _, status := range statuses {
		switch status.TbName {
		case "jp_5d_mva_ranking":
			jp5dMvaRankingDate = status.Date
		case "jp_daily_price":
			jpDailyPriceDate = status.Date
		case "jp_stocks_info":
			jpStocksInfoDate = status.Date
		}
	}
	fmt.Println("Out getLatestUpdateDates")
	return jp5dMvaRankingDate, jpDailyPriceDate, jpStocksInfoDate
}

// fetchRankingData はランキングデータを取得し、DTO にマッピングして戻す
func (s *RankingServiceImpl) fetchRankingData() (*[]dto.RankingServiceResponse, error) {
	// ランキングデータを取得
	rankingData, err := s.j5mrRepo.Get5dMvaRankingData()
	if err != nil {
		return nil, fmt.Errorf("failed to get ranking data: %w", err)
	}

	// ランキングデータからティッカーのリストを作成
	tickersSet := make(map[string]struct{})
	for _, data := range *rankingData {
		tickersSet[data.Symbol] = struct{}{}
	}
	var tickers []string
	for ticker := range tickersSet {
		tickers = append(tickers, ticker)
	}

	// ティッカーに対応する銘柄情報を取得
	stockInfoMap, err := s.jsiRepo.GetStockInfoByTickers(tickers)
	if err != nil {
		return nil, fmt.Errorf("failed to get stock info: %w", err)
	}

	// ティッカーに対応する最新の終値を取得
	latestPrices, err := s.jdpRepo.GetLatestClosePricesByTickers(tickers)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest close prices: %w", err)
	}

	// DTO に変換
	response := make([]dto.RankingServiceResponse, 0, len(*rankingData))
	for _, data := range *rankingData {
		stockInfo, ok := stockInfoMap[data.Symbol]
		if !ok {
			return nil, fmt.Errorf("stock info not found for ticker %s", data.Symbol)
		}

		latestPrice, ok := latestPrices[data.Symbol]
		if !ok {
			return nil, fmt.Errorf("latest price not found for ticker %s", data.Symbol)
		}

		response = append(response, dto.RankingServiceResponse{
			Ranking:     data.Ranking,
			Ticker:      data.Symbol,
			Date:        data.Date.Format("2006-01-02"),
			AvgTurnover: data.AvgTurnover,
			Name:        stockInfo.Name,
			LatestClose: latestPrice,
		})
	}

	return &response, nil
}
