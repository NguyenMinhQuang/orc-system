package config

import (
	"errors"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppVersion  string `envconfig:"APP_VERSION"`
	Port        string `envconfig:"PORT"`
	IsDebug     bool   `envconfig:"IS_DEBUG"`
	Stage       string `envconfig:"STAGE"`
	ServiceHost string `envconfig:"SERVICE_HOST"`
	SSL         bool   `envconfig:"SSL"`

	Mysql struct {
		Host       string `envconfig:"DB_HOST"`
		ReaderHost string `envconfig:"DB_READER_HOST"`
		UserName   string `envconfig:"DB_USERNAME"`
		PassWord   string `envconfig:"DB_PASSWORD"`
		DBName     string `envconfig:"DB_NAME"`
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

func LoadConfig() (*Config, error) {
	var config = &Config{}
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	err = envconfig.Process("", config)
	if err != nil {
		return nil, errors.New("failed to decode config env")
	}
	return config, nil
}
