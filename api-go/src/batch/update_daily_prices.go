// api-go\src\batch\update_daily_prices.go
package batch

import (
	"api-go/src/model"
	"api-go/src/repository"
	client "api-go/src/service/ms_gateway/client"
	gsdwd "api-go/src/service/ms_gateway/get_stocks_datalist_with_dates"
	"context"
	"fmt"
	"time"
	// その他必要なインポート
)

// UpdateDailyPrices は日次株価データを更新し、ステータスを更新します
func UpdateDailyPrices(ctx context.Context, udsRepo repository.UpdateStatusRepository, jsiRepo repository.JpStockInfoRepository, jdpRepo repository.JpDailyPriceRepository, clients map[string]interface{}, startDate, endDate string) error {
	gsdwdClient, ok := clients["get_stocks_datalist_with_dates"].(client.GetStocksDatalistWithDatesClient)
	if !ok {
		return fmt.Errorf("failed to get get_stocks_datalist_with_dates_client")
	}

	var Symbols []string
	stocks, err := jsiRepo.GetAllStockInfo()
	if err != nil {
		return fmt.Errorf("failed to get all stock info: %w", err)
	}

	for _, stock := range *stocks {
		symbol := stock.Ticker + ".T"
		Symbols = append(Symbols, symbol)
	}

	req := &gsdwd.GetStocksDatalistWithDatesRequest{
		Symbols:   Symbols,
		StartDate: startDate,
		EndDate:   endDate,
	}

	gsdwdResponse, err := gsdwdClient.GetStocksDatalist(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get stocks data list with dates: %w", err)
	}

	var newDailyPrices []model.JpDailyPrice
	for _, data := range gsdwdResponse.StockPrices {
		date, err := time.Parse("2006-01-02", data.Date)
		if err != nil {
			return fmt.Errorf("failed to parse date: %w", err)
		}
		dp := model.JpDailyPrice{
			Ticker:   data.Symbol,
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

	if err := jdpRepo.AddDailyPriceData(&newDailyPrices); err != nil {
		return fmt.Errorf("failed to add daily price data: %w", err)
	}

	if err := udsRepo.UpdateStatus("jp_daily_price"); err != nil {
		return fmt.Errorf("failed to update status for jp_daily_price: %w", err)
	}

	return nil
}
