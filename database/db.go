package database

import (
	"github.com/sajanjswl/auth/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb(logger *zap.Logger, cfg *config.Config) *gorm.DB {
	var err error
	dns := cfg.DBUserName + ":" + cfg.DBPassword + "@tcp" + "(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName + "?" + "charset=utf8mb4&parseTime=True&loc=Local"
	logger.Info("dns config", zap.String("dns", dns))
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		logger.Fatal("error connecting to database", zap.Error(err))
		return nil
	}

	return db
}
