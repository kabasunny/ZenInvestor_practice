// api-go\src\src\batch\update_ranking.go

package batch

import (
	"api-go/src/repository"
	"context"
	"fmt"
	// その他必要なインポート
)

// UpdateRanking は5日間平均ランキングデータを計算し、ステータスを更新
func UpdateRanking(ctx context.Context, udsRepo repository.UpdateStatusRepository, j5mrRepo repository.JP5dMvaRankingRepository) error {
	if err := j5mrRepo.Add5dMvaRankingData(); err != nil {
		return fmt.Errorf("failed to add 5d mva ranking data: %w", err)
	}

	if err := udsRepo.UpdateStatus("jp_5d_mva_ranking"); err != nil {
		return fmt.Errorf("failed to update status for jp_5d_mva_ranking: %w", err)
	}

	return nil
}
