package main

import (
	"log"
	"orc-system/config"
	"orc-system/pkg/logger"
)

func main() {
	log.Println("Starting api server")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	appLogger := logger.NewAppLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Stage: %s, SSL: %v", cfg.AppVersion, cfg.Logger.Level, cfg.Stage, cfg.SSL)
}
