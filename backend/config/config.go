package config

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func GetDatabaseConfig() (*DbConfig, error){
	err := godotenv.Load("../.env")
	if err != nil {
		zap.L().Error("Failed to load .env file", zap.Error(err))
		return nil, err
	}

	cfg := &DbConfig{
		Host:     os.Getenv("host"),
		Port:     os.Getenv("port"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Dbname:   os.Getenv("POSTGRES_DB"),
	}

	return cfg, nil
}
