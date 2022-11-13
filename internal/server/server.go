package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"orc-system/config"
	"orc-system/internal/middleware"
	"orc-system/pkg/logger"
	"os"
	"os/signal"
	"time"
)

// Server struct
type Server struct {
	echo       *echo.Echo
	tokenMaker middleware.Maker
	cfg        *config.Config
	db         *gorm.DB
}

func NewServer(cfg *config.Config, db *gorm.DB) (*Server, error) {
	tokenMaker, err := middleware.NewJWTMaker(cfg.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	return &Server{
		echo:       echo.New(),
		tokenMaker: tokenMaker,
		cfg:        cfg,
		db:         db}, nil
}

func (s *Server) Run() error {
	// // Setup
	if err := s.NewHTTPHandler(s.echo); err != nil {
		return err
	}

	// Start server
	go func() {
		if err := s.echo.Start(":" + s.cfg.Port); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Error starting Server: ", err)
		}
	}()

	// gracefull Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.echo.Server.Shutdown(ctx)
}
