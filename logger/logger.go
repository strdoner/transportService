package logger

import (
	"go.uber.org/zap"
	"os"
)

func Init() *zap.Logger {
	logger, _ := zap.NewProduction()
	if os.Getenv("ENV") == "local" {
		logger, _ = zap.NewDevelopment()
	}
	zap.ReplaceGlobals(logger)
	return logger
}
