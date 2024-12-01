// api-go\src\service\ranking_service.go

package service

import (
	"api-go/src/dto"
)

// RankingService は株価データを取得するためのインターフェース
type RankingService interface {

	// 全ランキングを取得
	// GetFullRankingData(ctx context.Context) (*[]dto.RankingServiceResponse, error)

	// 100位までのランキングを取得
	GetTop100RankingData() (*[]dto.RankingServiceResponse, error)

	// 範囲を指定してランキングを取得
	GetRankingDataByRange(startRank int, endRank int) (*[]dto.RankingServiceResponse, error)
}
