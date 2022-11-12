package delivery

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"orc-system/config"
	"orc-system/internal/token"
	"orc-system/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	certFile       = "ssl/Server.crt"
	keyFile        = "ssl/Server.pem"
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

// Server struct
type Server struct {
	echo       *echo.Echo
	tokenMaker token.Maker
	cfg        *config.Config
	db         *gorm.DB
	logger     logger.Logger
}

func NewServer(cfg *config.Config, db *gorm.DB, logger logger.Logger) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(cfg.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	return &Server{
		echo:       echo.New(),
		tokenMaker: tokenMaker,
		cfg:        cfg,
		db:         db,
		logger:     logger}, nil
}

func (s *Server) Run() error {

	sv := &http.Server{
		Addr:           s.cfg.Port,
		MaxHeaderBytes: maxHeaderBytes,
	}
	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Port)
		if err := s.echo.StartServer(sv); err != nil {
			s.logger.Fatalf("Error starting Server: ", err)
		}
	}()

	if err := s.NewHTTPHandler(s.echo); err != nil {
		return err
	}

	// gracefull Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()
	s.logger.Info("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}
