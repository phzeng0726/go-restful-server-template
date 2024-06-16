package database

import (
	"log"

	"github.com/phzeng0726/go-server-template/internal/config"
	"github.com/phzeng0726/go-server-template/internal/domain"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(logger *zap.Logger) *gorm.DB {
	conn, err := gorm.Open(postgres.Open(config.Env.DatabaseDSN), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect database", zap.Error(err))
	}

	logger.Info("Database connected")
	return conn
}

// 確保DB和Model的格式對的上
func SyncDatabase(conn *gorm.DB) {
	err := conn.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("Failed to migrate User: %v", err)
	}
}
