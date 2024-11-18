// api-go\src\service\ranking_service.go

package service

import (
	"api-go/src/dto"
	"context"
)

// RankingService は株価データを取得するためのインターフェース
type RankingService interface {
	GetRankingData(ctx context.Context) (dto.RankingServiceRespons, error)
}
