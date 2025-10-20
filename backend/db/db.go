package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"transportService/config"
)

func StartSQLConnection() (*sql.DB, error) {
	cfg, err := config.GetDatabaseConfig()
	if err != nil {
		return nil, err
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	zap.L().Info("Database connected successfully",
		zap.String("host", cfg.Host),
		zap.String("port", cfg.Port),
	)

	return db, nil
}
