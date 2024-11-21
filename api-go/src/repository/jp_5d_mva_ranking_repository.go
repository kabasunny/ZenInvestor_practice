// api-go\src\repository\jp_5d_mva_ranking_repository.go
package repository

import "api-go/src/model"

type JP5dMvaRankingRepository interface {
	// 売買代金5日平均ランキングデータを取得する:データ取得時(毎回)
	Get5dMvaRankingData() (*[]model.Jp5dMvaRanking, error)

	// 売買代金5日平均ランキングデータを追加する:一日一回
	// 引数は無し、jp_daily_priceテーブルからjp_daily_priceテーブルの最新日付から5日間分の日付をさかのぼり、ticker毎に5日間平均をjp_5d_mva_rankingテーブルに格納する
	Add5dMvaRankingData() error

	// 売買代金5日平均ランキングデータを削除する:データが一定数に達したら(一日一回)
	Delete5dMvaRankingData(days int) error

	// その他必要に応じて
}
