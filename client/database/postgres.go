package database

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"orc-system/config"
	"time"
)

var db *gorm.DB

func NewPsqlDB(cfg *config.Config) (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		cfg.Postgres.Host,
		cfg.Postgres.UserName,
		cfg.Postgres.PassWord,
		cfg.Postgres.DBName,
		cfg.Postgres.Port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db = db.Debug()
	connection, err := db.DB()
	if err != nil {
		return nil, err
	}

	connection.SetMaxIdleConns(cfg.Postgres.DBMaxIdleConns)
	connection.SetMaxOpenConns(cfg.Postgres.DBMaxOpenConns)
	connection.SetConnMaxLifetime(time.Second * 14400)

	return db.Session(&gorm.Session{}), nil
}

func DisConnection() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	if db != nil {
		sqlDB.Close()
	}

}
