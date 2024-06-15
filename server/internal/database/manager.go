package database

import (
	"fmt"
	"log"
	"time"

	"github.com/phzeng0726/go-server-template/internal/domain"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(logger *zap.Logger) *gorm.DB {
	// https://github.com/go-gorm/postgres
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		logger.Fatal("Failed to connect database", zap.Error(err))
	}

	logger.Info("Sqlite database connected")

	return conn
}

// 確保DB和Model的格式對的上
func SyncDatabase(conn *gorm.DB) {
	err := conn.AutoMigrate(&domain.VacuumInfo{})
	if err != nil {
		log.Fatalf("Failed to migrate VacuumInfo: %v", err)
	}
}

// 釋放儲存空間
func ReleaseSpaceInDatabase(db *gorm.DB) {
	var vacuumInfo domain.VacuumInfo
	if err := db.FirstOrCreate(&vacuumInfo, domain.VacuumInfo{Id: 1}).Error; err != nil {
		fmt.Printf("Failed to get or create vacuum info: %v", err)
		return
	}

	currentTime := time.Now().Unix()
	timeDiff := currentTime - vacuumInfo.LastVacuumTime

	// 一周執行一次
	if timeDiff > 604800 {
		if err := db.Exec("VACUUM;").Error; err != nil {
			fmt.Printf("Failed to execute VACUUM: %v", err)
			return
		}
		vacuumInfo.LastVacuumTime = currentTime
	}

	if err := db.Save(&vacuumInfo).Error; err != nil {
		fmt.Printf("Failed to save vacuum info: %v", err)
		return
	}
}
