package database

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/yangliang4488/goblog/pkg/logger"
)

var DB *sql.DB

func Initialize() {
	InitDB()
	CreateTables()
}

func InitDB() {
	var err error
	config := mysql.Config{
		User:                 "root",
		Passwd:               "123456",
		Addr:                 "127.0.0.1:3308",
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}
	// 准备连接池
	DB, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)
	// 设置最大空闲连接数
	DB.SetMaxOpenConns(25)
	// 设置最大连接数
	DB.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	DB.SetConnMaxLifetime(5 * time.Minute)

	err = DB.Ping()
	logger.LogError(err)

}

func CreateTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
		id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
		body longtext COLLATE utf8mb4_unicode_ci);`

	_, err := DB.Exec(createArticlesSQL)
	logger.LogError(err)
}
