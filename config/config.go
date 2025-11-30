package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB 全局数据库连接池（必须在InitDB后初始化）
var DB *pgxpool.Pool

// InitDB 初始化PostgreSQL连接池
func InitDB() error {
	// 从环境变量或配置文件读取数据库连接信息（根据你的实际配置调整）
	connStr := "postgres://system:123456@localhost:54321/hotel_db?sslmode=disable"
	// （可选）从环境变量读取：connStr = os.Getenv("DATABASE_URL")

	// 创建连接池
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Printf("无法创建数据库连接池: %v", err)
		return err
	}

	// 测试连接
	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		log.Printf("数据库连接失败: %v", err)
		return err
	}

	DB = pool
	log.Println("数据库连接池初始化成功")
	return nil
}
