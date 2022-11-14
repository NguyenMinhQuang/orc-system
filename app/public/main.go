package main

import (
	"log"
	"orc-system/config"
	"orc-system/internal/server"
	"orc-system/pkg/database"
	"orc-system/pkg/logger"
)

func main() {
	log.Println("Starting api server")
	cfg := config.GetConfig()
	logger.SetLevel(cfg.Logger.Level)

	sqlDB, err := database.NewMysqlDB(cfg)
	if err != nil {
		logger.Fatalf("Mysql init: %s", err)
	}
	defer database.DisConnect()
	s, err := server.NewServer(cfg, sqlDB)
	if err != nil {
		log.Println("cannot create server:", err)
		return
	}
	if err = s.Run(); err != nil {
		log.Println(err)
		return
	}
}
