package main

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"transportService/handlers"
	"transportService/logger"
	"transportService/repository"
	"transportService/services"
)

func main() {
	logger.Init()
	//db, err := dbPackage.StartSQLConnection()
	//defer func(db *sql.DB) {
	//	err = db.Close()
	//	if err != nil {
	//		zap.L().Error("Error via closing database", zap.Error(err))
	//	}
	//}(db)
	//
	//if err != nil {
	//	zap.L().Error("Failed connecting to database", zap.Error(err))
	//	return
	//}

	stub := repository.NewParkingStub()

	service := services.NewParkingService(stub)

	parkingHandler := handlers.NewParkingHandler(service)
	mux := handlers.NewRouter(parkingHandler)

	server := services.NewServer(mux)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.L().Error("Server failed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	zap.L().Info("Shutting down...")
	if err := server.Stop(ctx); err != nil {
		zap.L().Error("Failed to shutdown server", zap.Error(err))
	}

}
