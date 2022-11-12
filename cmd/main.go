package main

import (
	"log"
	"orc-system/client/database"
	"orc-system/config"
	"orc-system/internal/server"
	"orc-system/pkg/logger"
)

func main() {
	log.Println("Starting api server")

	cfg := config.GetConfig()

	appLogger := logger.NewAppLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Stage: %s, SSL: %v", cfg.AppVersion, cfg.Logger.Level, cfg.Stage, cfg.SSL)

	sqlDB, err := database.NewMysqlDB(cfg)
	if err != nil {
		appLogger.Fatalf("Postgresql init: %s", err)
	}
	defer database.DisConnect()
	s, err := server.NewServer(cfg, sqlDB, appLogger)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
