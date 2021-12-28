package dbs

import (
	"fmt"

	"swift_typing_api/common"
	"swift_typing_api/conf"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDatabase interface {
	GetInstance() *gorm.DB
}

type database struct {
	db *gorm.DB
}

// NewDatabase return new IDatabase interface
func NewDatabase() IDatabase {
	dbConfig := conf.Config.Database
	connectionPath := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port, dbConfig.SSLMode)

	db, err := gorm.Open(postgres.Open(connectionPath), &gorm.Config{})
	if err != nil {
		common.GetLogger().Error("error")
		//logger.Fatal("Cannot connect to database: ", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		common.GetLogger().Error("error")
		//logger.Fatal("Cannot connect to database: ", err)
	}
	// Set up connection pool
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)
	return &database{
		db: db,
	}
}

// GetInstance get database instance
func (d *database) GetInstance() *gorm.DB {
	return d.db
}
