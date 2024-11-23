// api-go\src\service\ranking_service_impl.go

package service

import (
	"api-go/src/dto"
	"api-go/src/model"
	"api-go/src/repository"
	client "api-go/src/service/ms_gateway/client"
	get_stock_info_jq "api-go/src/service/ms_gateway/get_stock_info_jq"
	gsdwd "api-go/src/service/ms_gateway/get_stocks_datalist_with_dates"
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
	if err != nil {
		return nil, fmt.Errorf("failed to get update statuses: %w", err)
	}

	// 今日の日付を取得
	today := time.Now().UTC().Truncate(24 * time.Hour)
	fmt.Println("Got Tody")

	// 各テーブルの更新日を格納する変数を初期化
	jp5dMvaRankingDate, jpDailyPriceDate, jpStocksInfoDate := getLatestUpdateDates(statuses)

	// 状態に応じたデータの更新と取得
	switch {
	// ケース 1: jp_5d_mva_ranking が最新の場合
	// そのままランキングデータを取得し、 DTO にマッピングして戻す処理を呼び出す
	case jp5dMvaRankingDate.Equal(today):
		fmt.Println("In Case1")
		return s.fetchRankingData()

	// ケース 2: jp_daily_price と jp_stocks_info が最新の場合
	// 5 日間平均ランキングデータを計算・更新し、更新日付を更新
	case jpDailyPriceDate.Equal(today) && jpStocksInfoDate.Equal(today):
		fmt.Println("In Case2")
		if err := s.updateRanking(ctx); err != nil {
			return nil, err
		}

	// ケース 3: jp_daily_price が古いが jp_stocks_info が最新の場合
	// 最新の株価データを取得・追加し、5 日間平均ランキングデータを計算・更新し、更新日付を更新
	case !jpDailyPriceDate.Equal(today) && jpStocksInfoDate.Equal(today):
		fmt.Println("In Case3")
		startDate := jpDailyPriceDate.AddDate(0, 0, 1).Format("2006-01-02")
		endDate := today.Format("2006-01-02")
		if err := s.updateDailyPrices(ctx, startDate, endDate); err != nil {
			return nil, err
		}
		if err := s.updateRanking(ctx); err != nil {
			return nil, err
		}

	// ケース 4: jp_daily_price と jp_stocks_info が古い場合
	// 銘柄情報と株価データの両方を取得・更新し、5 日間平均ランキングデータを計算・更新し、更新日付を更新
	case !jpDailyPriceDate.Equal(today) && !jpStocksInfoDate.Equal(today):
		fmt.Println("In Case4")
		if err := s.updateStockInfo(ctx); err != nil {
			return nil, err
		}
		startDate := jpDailyPriceDate.AddDate(0, 0, 1).Format("2006-01-02")
		endDate := today.Format("2006-01-02")
		if err := s.updateDailyPrices(ctx, startDate, endDate); err != nil {
			return nil, err
		}
		if err := s.updateRanking(ctx); err != nil {
			return nil, err
		}
	}

	return s.fetchRankingData()
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
	return jp5dMvaRankingDate, jpDailyPriceDate, jpStocksInfoDate
}

// fetchRankingData はランキングデータを取得し、DTO にマッピングして戻す
func (s *RankingServiceImpl) fetchRankingData() (*[]dto.RankingServiceResponse, error) {
	rankingData, err := s.j5mrRepo.Get5dMvaRankingData()
	if err != nil {
		return nil, fmt.Errorf("failed to get ranking data: %w", err)
	}
	return s.mapRankingDataToResponse(rankingData)
}

// updateRanking は5日間平均ランキングデータを計算し、ステータスを更新
func (s *RankingServiceImpl) updateRanking(ctx context.Context) error {
	if err := s.j5mrRepo.Add5dMvaRankingData(); err != nil {
		return fmt.Errorf("failed to add 5d mva ranking data: %w", err)
	}

	if err := s.udsRepo.UpdateStatus("jp_5d_mva_ranking"); err != nil {
		return fmt.Errorf("failed to update status for jp_5d_mva_ranking: %w", err)
	}

	return nil
}

// updateDailyPrices は日次株価データを更新し、ステータスを更新します
func (s *RankingServiceImpl) updateDailyPrices(ctx context.Context, startDate, endDate string) error {
	gsdwdClient, ok := s.clients["get_stocks_datalist_with_dates"].(client.GetStocksDatalistWithDatesClient)
	if !ok {
		return fmt.Errorf("failed to get get_stocks_datalist_with_dates_client")
	}

	req := &gsdwd.GetStocksDatalistWithDatesRequest{
		StartDate: startDate,
		EndDate:   endDate,
	}

	gsdwdResponse, err := gsdwdClient.GetStocksDatalist(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get stocks data list with dates: %w", err)
	}

	var newDailyPrices []model.JpDailyPrice
	for _, data := range gsdwdResponse.StockPrices { // 修正ポイント: Data から StockPrices へ　// protoファイルではstock_prices
		date, err := time.Parse("2006-01-02", data.Date)
		if err != nil {
			return fmt.Errorf("failed to parse date: %w", err)
		}
		dp := model.JpDailyPrice{
			Ticker:   data.Symbol, // 修正ポイント: data.Ticker から data.Symbol へ
			Date:     date,
			Open:     data.Open,
			Close:    data.Close,
			High:     data.High,
			Low:      data.Low,
			Volume:   data.Volume,
			Turnover: data.Turnover,
		}
		newDailyPrices = append(newDailyPrices, dp)
	}

	if err := s.jdpRepo.AddDailyPriceData(&newDailyPrices); err != nil {
		return fmt.Errorf("failed to add daily price data: %w", err)
	}

	if err := s.udsRepo.UpdateStatus("jp_daily_price"); err != nil {
		return fmt.Errorf("failed to update status for jp_daily_price: %w", err)
	}

	return nil
}

// updateStockInfo は銘柄情報を更新し、ステータスを更新
func (s *RankingServiceImpl) updateStockInfo(ctx context.Context) error {
	fmt.Println("In updateStockInfo")
	stockInfoClient, ok := s.clients["get_stock_info_jq"].(client.GetStockInfoJqClient)
	if !ok {
		return fmt.Errorf("failed to get get_stock_info_jq_client")
	}

	stockInfoReq := &get_stock_info_jq.GetStockInfoJqRequest{}

	stockInfoRes, err := stockInfoClient.GetStockInfoJq(ctx, stockInfoReq)
	if err != nil {
		return fmt.Errorf("failed to get stock info: %w", err)
	}

	var newStockInfos []model.JpStockInfo
	for _, data := range stockInfoRes.Stocks { // 修正ポイント: Data から Stocks へ　// protoファイルではstocks
		si := model.JpStockInfo{
			Ticker:   data.Ticker,
			Name:     data.Name,
			Sector:   data.Sector,
			Industry: data.Industry,
		}
		newStockInfos = append(newStockInfos, si)
	}

	if err := s.jsiRepo.UpdateStockInfo(&newStockInfos); err != nil {
		return fmt.Errorf("failed to update stock info: %w", err)
	}

	if err := s.udsRepo.UpdateStatus("jp_stocks_info"); err != nil {
		return fmt.Errorf("failed to update status for jp_stocks_info: %w", err)
	}

	return nil
}

// mapRankingDataToResponse はランキングデータを DTO にマッピング
func (s *RankingServiceImpl) mapRankingDataToResponse(rankingData *[]model.Jp5dMvaRanking) (*[]dto.RankingServiceResponse, error) {
	// ランキングデータからティッカーのリストを作成
	tickersSet := make(map[string]struct{})
	for _, data := range *rankingData {
		tickersSet[data.Ticker] = struct{}{}
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
		stockInfo, ok := stockInfoMap[data.Ticker]
		if !ok {
			return nil, fmt.Errorf("stock info not found for ticker %s", data.Ticker)
		}

		latestPrice, ok := latestPrices[data.Ticker]
		if !ok {
			return nil, fmt.Errorf("latest price not found for ticker %s", data.Ticker)
		}

		response = append(response, dto.RankingServiceResponse{
			Ranking:     data.Ranking,
			Ticker:      data.Ticker,
			Date:        data.Date.Format("2006-01-02"),
			AvgTurnover: data.AvgTurnover,
			Name:        stockInfo.Name,
			LatestClose: latestPrice,
		})
	}

	return &response, nil
}
