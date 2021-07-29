package sqlite

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	log.Println("[InitDB] init sqlite connection...")

	gormDB, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Panicln("failed to connect sqlite")
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Panicln("failed to get sqlDB")
		return nil, err
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	// set debug to inspect SQL
	gormDB = gormDB.Debug()
	return gormDB, nil
}
