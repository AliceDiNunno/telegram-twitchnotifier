package postgres

import (
	"TwitchNotifierForTelegram/src/config"
	"fmt"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func StartGormDatabase(config *config.Database) *gorm.DB {
	if config == nil {
		panic("Database config is nil")
		return nil
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)
	db, err := gorm.Open(pg.Open(psqlInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}
	db.Logger.LogMode(logger.Info)
	return db
}
