package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"orc-system/config"
	"time"
)

var db *gorm.DB

func NewMysqlDB(cfg *config.Config) (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Mysql.UserName,
		cfg.Mysql.PassWord,
		cfg.Mysql.Host,
		cfg.Mysql.Port,
		cfg.Mysql.DBName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if cfg.Mysql.DBDebugMode {
		db = db.Debug()
	}

	connection, err := db.DB()
	if err != nil {
		return nil, err
	}

	connection.SetMaxIdleConns(cfg.Mysql.DBMaxIdleConns)
	connection.SetMaxOpenConns(cfg.Mysql.DBMaxOpenConns)
	connection.SetConnMaxLifetime(time.Second * 14400)

	err = connection.Ping()
	if err != nil {
		return nil, err
	}
	return db.Session(&gorm.Session{}), nil
}

func DisConnect() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	if db != nil {
		sqlDB.Close()
	}

}
