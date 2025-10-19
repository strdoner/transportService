package config

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func GetDatabaseConfig( (*DbConfig, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		zap.L().Error("Failed to load .env file", zap.Error(err))
		return nil, err
	}

	cfg := &DbConfig{
		Host:     os.Getenv("host"),
		Port:     os.Getenv("port"),
		User:     os.Getenv("user"),
		Password: os.Getenv("password"),
		Dbname:   os.Getenv("dbname"),
	}

	return cfg, nil
}
