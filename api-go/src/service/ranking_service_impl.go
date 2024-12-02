// api-go\src\service\ranking_service_impl.go

package service

import (
	"api-go/src/dto"
	"api-go/src/model"
	"api-go/src/repository"
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

// GetRankingDataByRange は指定した順位の範囲でランキングデータを取得
func (s *RankingServiceImpl) GetRankingDataByRange(startRank int, endRank int) (*[]dto.RankingServiceResponse, error) {
	fmt.Println("In GetRankingDataByRange")
	// update_statusテーブルの構造体を取得
	statuses, err := s.udsRepo.GetAllUpdateStatuses()
	if err != nil {
		return nil, fmt.Errorf("failed to get update statuses: %w", err)
	}

	// 今日の日付を取得
	today := time.Now()
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())

	// 各テーブルの更新日を格納する変数を初期化
	jp5dMvaRankingDate, _, _ := getLatestUpdateDates(statuses)

	// 状態に応じたデータの取得
	if !jp5dMvaRankingDate.Equal(today) {
		return nil, fmt.Errorf("data is not up-to-date. Please run batch processes")
	}

	// ランキングデータを取得
	rankingData, err := s.j5mrRepo.Get5dMvaRankingData()
	if err != nil {
		return nil, fmt.Errorf("failed to get ranking data: %w", err)
	}

	// 指定された範囲のランキングデータを抽出
	if startRank < 1 || endRank > len(*rankingData) || startRank > endRank {
		return nil, fmt.Errorf("invalid rank range specified")
	}

	rankingDataByRange := (*rankingData)[startRank-1 : endRank]

	// ランキングデータからティッカーのリストを作成
	symbolsSet := make(map[string]struct{})
	for _, data := range rankingDataByRange {
		symbolsSet[data.Symbol] = struct{}{}
	}
	var symbols []string
	for symbol := range symbolsSet {
		symbols = append(symbols, symbol)
	}

	// ティッカーに対応する銘柄情報を取得
	stockInfoMap, err := s.jsiRepo.GetStockInfoByTickers(symbols)
	if err != nil {
		return nil, fmt.Errorf("failed to get stock info: %w", err)
	}

	// ティッカーに対応する最新の終値を取得
	latestPrices, err := s.jdpRepo.GetLatestClosePricesByTickers(symbols)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest close prices: %w", err)
	}

	// DTO に変換
	response := make([]dto.RankingServiceResponse, 0, len(rankingDataByRange))
	for _, data := range rankingDataByRange {
		stockInfo, ok := stockInfoMap[data.Symbol]
		if !ok {
			return nil, fmt.Errorf("stock info not found for symbol %s", data.Symbol)
		}

		latestPrice, ok := latestPrices[data.Symbol]
		if !ok {
			return nil, fmt.Errorf("latest price not found for symbol %s", data.Symbol)
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

// GetTop100RankingData は上位100銘柄のランキングデータを取得 UIから初回取得用
func (s *RankingServiceImpl) GetTop100RankingData() (*[]dto.RankingServiceResponse, error) {
	return s.GetRankingDataByRange(1, 100)
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
