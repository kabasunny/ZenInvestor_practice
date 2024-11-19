// api-go\src\repository\jp_mva5d_ranking_repository.go
package repository

type JP5dMvaRankingRepository interface {
	// 売買代金5日平均ランキングデータを取得する:データ取得時(毎回)
	Get5dMvaRankingData()

	// 売買代金5日平均ランキングデータを追加する:一日一回
	Add5dMvaRankingData()

	// 売買代金5日平均ランキングデータを削除する:データが一定数に達したら(一日一回)
	Delete5dMvaRankingData()
}
