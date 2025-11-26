package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func GetDatabaseConfig() (*DbConfig, error) {
	// Загружаем .env (игнорируем ошибку, если файла нет)
	_ = godotenv.Load("../.env")

	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	if user == "" || password == "" || dbname == "" {
		return nil, fmt.Errorf("missing DB credentials in .env")
	}

	return &DbConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Dbname:   dbname,
	}, nil
}
