// api-go\src\repository\jp_5d_mva_ranking_repository_impl.go

package repository

import (
	"api-go/src/model"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type jp5dMvaRankingRepositoryImpl struct {
	db *gorm.DB
}

func NewJp5dMvaRankingRepository(db *gorm.DB) JP5dMvaRankingRepository {
	return &jp5dMvaRankingRepositoryImpl{db: db}
}

// 売買代金5日平均ランキングデータを取得する（最新の日付のみ）
func (r *jp5dMvaRankingRepositoryImpl) Get5dMvaRankingData() (*[]model.Jp5dMvaRanking, error) {
	var rankings []model.Jp5dMvaRanking

	// サブクエリで最新の日付を取得
	var latestDate time.Time
	if err := r.db.Model(&model.Jp5dMvaRanking{}).Select("MAX(date)").Scan(&latestDate).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch the latest date: %w", err)
	}

	// 最新の日付のデータを取得
	if err := r.db.Where("date = ?", latestDate).Find(&rankings).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch 5d MVA ranking data: %w", err)
	}

	return &rankings, nil
}

// 売買代金5日平均ランキングデータを追加する
func (repo *jp5dMvaRankingRepositoryImpl) Add5dMvaRankingData() error {
	fmt.Println("In Add5dMvaRankingData")
	startTimeOverall := time.Now()

	// 最新日付を含む直近5営業日を取得
	var tradingDates []time.Time
	err := repo.db.Model(&model.JpDailyPrice{}).
		Distinct("date").
		Order("date DESC").
		Limit(5).
		Pluck("date", &tradingDates).Error
	if err != nil {
		return err
	}

	if len(tradingDates) < 5 {
		return fmt.Errorf("5日分のデータがありません")
	}

	latestDate := tradingDates[0]

	// 5日間平均を計算し、ランキングを生成
	rows, err := repo.db.Raw(
		`
		SELECT 
		symbol, 
		CEIL(AVG(turnover)) as avg_turnover,
		RANK() OVER (ORDER BY AVG(turnover) DESC) as ranking,
		? AS date
		FROM jp_daily_price
		WHERE date IN (?)
		GROUP BY symbol
		`, latestDate, tradingDates).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	// トランザクション開始
	tx := repo.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for rows.Next() {
		var ranking model.Jp5dMvaRanking
		// ドライバが日付データを []uint8 として返すため、time.Time に直接スキャンできない
		// 文字列として取得し、time.Parse で変換する
		var dateStr string
		if err := rows.Scan(&ranking.Symbol, &ranking.AvgTurnover, &ranking.Ranking, &dateStr); err != nil {
			tx.Rollback()
			return err
		}

		// 文字列からtime.Time型に変換
		date, err := time.Parse("2006-01-02", dateStr) // dateStrを解析してtime.Time型に変換
		if err != nil {
			tx.Rollback()
			return err
		}
		ranking.Date = date // 変換されたtime.Time型をranking.Dateにセット

		err = tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "ranking"}, {Name: "symbol"}, {Name: "date"}},
			DoUpdates: clause.AssignmentColumns([]string{"avg_turnover"}),
		}).Create(&ranking).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := rows.Err(); err != nil {
		tx.Rollback()
		return err
	}

	endTimeOverall := time.Now()
	fmt.Printf("raking作成のバッチの処理時間: %s\n", endTimeOverall.Sub(startTimeOverall))
	fmt.Println("Out Add5dMvaRankingData")
	return tx.Commit().Error
}

// 売買代金5日平均ランキングデータを削除する
func (r *jp5dMvaRankingRepositoryImpl) Delete5dMvaRankingData(days int) error {
	// 現在の日付から指定された日数を引いた日付を計算
	beforeDate := time.Now().AddDate(0, 0, -days)

	// 指定された日付以前のデータを削除
	result := r.db.Where("date < ?", beforeDate).Delete(&model.Jp5dMvaRanking{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete 5d MVA ranking data: %w", result.Error)
	}

	fmt.Printf("Deleted %d records older than %s\n", result.RowsAffected, beforeDate)
	return nil
}
