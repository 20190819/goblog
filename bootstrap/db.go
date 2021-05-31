package bootstrap

import (
	"time"

	"github.com/yangliang4488/goblog/app/models/article"
	"github.com/yangliang4488/goblog/app/models/user"
	"github.com/yangliang4488/goblog/pkg/model"
	"gorm.io/gorm"
)

const MAX_OPEN = 100
const MAX_IDLE = 25
const MAX_LIFETIME = 5 * time.Minute

func SetupDB() {
	db := model.ConnectDB()
	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(MAX_OPEN)
	sqlDB.SetMaxIdleConns(MAX_IDLE)
	sqlDB.SetConnMaxLifetime(MAX_LIFETIME)

	CusorMigration(db)
}

func CusorMigration(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
	)
}
