// api-go\src\batch\update_stock_info.go

package batch

import (
	"api-go/src/model"
	"api-go/src/repository"
	client "api-go/src/service/ms_gateway/client"
	get_stock_info_jq "api-go/src/service/ms_gateway/get_stock_info_jq"
	"context"
	"fmt"
	// その他必要なインポート
)

// UpdateStockInfo は銘柄情報を更新し、ステータスを更新します
func UpdateStockInfo(ctx context.Context, udsRepo repository.UpdateStatusRepository, jsiRepo repository.JpStockInfoRepository, clients map[string]interface{}) error {
	stockInfoClient, ok := clients["get_stock_info_jq"].(client.GetStockInfoJqClient)
	if !ok {
		return fmt.Errorf("failed to get get_stock_info_jq_client")
	}

	stockInfoReq := &get_stock_info_jq.GetStockInfoJqRequest{}

	stockInfoRes, err := stockInfoClient.GetStockInfoJq(ctx, stockInfoReq)
	if err != nil {
		return fmt.Errorf("failed to get stock info: %w", err)
	}

	var newStockInfos []model.JpStockInfo
	for _, data := range stockInfoRes.Stocks {
		si := model.JpStockInfo{
			Ticker:   data.Ticker,
			Name:     data.Name,
			Sector:   data.Sector,
			Industry: data.Industry,
		}
		newStockInfos = append(newStockInfos, si)
	}

	if err := jsiRepo.UpdateStockInfo(&newStockInfos); err != nil {
		return fmt.Errorf("failed to update stock info: %w", err)
	}

	if err := udsRepo.UpdateStatus("jp_stocks_info"); err != nil {
		return fmt.Errorf("failed to update status for jp_stocks_info: %w", err)
	}

	return nil
}
