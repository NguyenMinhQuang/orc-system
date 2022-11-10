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

	psqlDB, err := database.NewPsqlDB(cfg)
	if err != nil {
		appLogger.Fatalf("Postgresql init: %s", err)
	}
	defer database.DisConnection()
	s := server.NewServer(cfg, psqlDB, appLogger)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
