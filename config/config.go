package config

import (
	"errors"
	"log"
	"orc-system/codetype"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var config *Config

type Config struct {
	AppVersion        string             `envconfig:"APP_VERSION"`
	Port              string             `envconfig:"PORT"`
	IsDebug           bool               `envconfig:"IS_DEBUG"`
	Stage             codetype.StageType `envconfig:"STAGE"`
	ServiceHost       string             `envconfig:"SERVICE_HOST"`
	SSL               bool               `envconfig:"SSL"`
	AllowOrigins      string             `envconfig:"ALLOW_ORIGINS"`
	TokenSymmetricKey string             `envconfig:"TOKEN_SYMMETRIC_KEY"`
	EndPoint          string             `envconfig:"END_POINT"`

	Postgres struct {
		Host           string `envconfig:"DB_HOST"`
		Port           string `envconfig:"DB_PORT"`
		ReaderHost     string `envconfig:"DB_READER_HOST"`
		UserName       string `envconfig:"DB_USERNAME"`
		PassWord       string `envconfig:"DB_PASSWORD"`
		DBName         string `envconfig:"DB_NAME"`
		DBMaxIdleConns int    `envconfig:"DB_MAX_IDLE_CONNS"`
		DBMaxOpenConns int    `envconfig:"DB_MAX_OPEN_CONNS"`
	}
	HealthCheck struct {
		HealthCheckEndPoint string `envconfig:"HEAL_CHECK_ENPOINT"`
	}

	Logger struct {
		DevMode           bool   `envconfig:"DEV_MODE"`
		DisableCaller     bool   `envconfig:"DISABLECALLER"`
		DisableStacktrace bool   `envconfig:"DISABLESTACKTRACE"`
		Encoding          string `envconfig:"ENCODING"`
		Level             string `envconfig:"LEVEL"`
	}
}

func init() {
	config = &Config{}
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	err = envconfig.Process("", config)
	if err != nil {
		log.Fatal(errors.New("failed to decode config env"))
	}

	config.Stage.UpCase()
}

func GetConfig() *Config {
	return config
}
