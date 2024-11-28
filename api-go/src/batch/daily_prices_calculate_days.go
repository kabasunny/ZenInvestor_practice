// api-go\src\batch\daily_prices_calculate_days.go

package batch

import (
	"api-go/src/repository"
	"context"
	"fmt"
	"time"
)

func CalculateLookbackDate(ctx context.Context, jdpRepo repository.JpDailyPriceRepository, startDate string, lookbackDays int) (int, error) {
	latestDate, err := jdpRepo.GetLatestDate() // DB内の銘柄情報として保持する日付を取得 トヨタ自動車 (7203)、ソニー (6758)、ソフトバンク (9984)のいずれか
	if err != nil {
		return 0, err
	}

	// startDate と latestDate を time.Time 型に変換
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return 0, fmt.Errorf("failed to parse startDate: %w", err)
	}

	latest, err := time.Parse("2006-01-02", latestDate)
	if err != nil {
		return 0, fmt.Errorf("failed to parse latestDate: %w", err)
	}

	// 日付の差を計算
	dateDifference := int(start.Sub(latest).Hours() / 24)

	// lookbackDays を調整
	trueLookbackDays := lookbackDays
	if dateDifference < lookbackDays {
		trueLookbackDays = lookbackDays - dateDifference
	}

	return trueLookbackDays, nil
}
