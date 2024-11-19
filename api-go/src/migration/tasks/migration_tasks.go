// api-go\src\migration\migration.go
package migration

import "gorm.io/gorm"

type MigrationTasks struct {
	Up   func(db *gorm.DB) error // データベースに変更を適用するための関数:現状テーブル更新
	Down func(db *gorm.DB) error // データベースの変更を元に戻すための関数:現状テーブル削除
}

var migrations = []MigrationTasks{
	// マスタテーブル
	{Up: Up_001, Down: Down_001},
	{Up: Up_002, Down: Down_002},

	// フラグテーブル
	{Up: Up_011, Down: Down_011},

	// 生成テーブル
	{Up: Up_021, Down: Down_021},
}

func MigrateUp(db *gorm.DB) error {
	for _, m := range migrations {
		if err := m.Up(db); err != nil {
			return err
		}
	}
	return nil
}

func MigrateDown(db *gorm.DB) error {
	for i := len(migrations) - 1; i >= 0; i-- {
		if err := migrations[i].Down(db); err != nil {
			return err
		}
	}
	return nil
}
