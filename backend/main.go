package main

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"transportService/db"
	"transportService/handlers"
	"transportService/logger"
	"transportService/repository"
	"transportService/services"

	"go.uber.org/zap"
)

func main() {
	logger.Init()
	sqlConnection, err := db.StartSQLConnection()
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			zap.L().Error("Error via closing database", zap.Error(err))
		}
	}(sqlConnection)

	if err != nil {
		zap.L().Error("Failed connecting to database", zap.Error(err))
		return
	}

	parkingRepo := repository.NewParkingRepository(sqlConnection)
	reservationRepo := repository.NewReservationRepository(sqlConnection)
	vehicleRepo := repository.NewVehicleRepository(sqlConnection)

	parkingService := services.NewParkingService(parkingRepo, reservationRepo)
	vehicleService := services.NewVehicleService(vehicleRepo)

	parkingHandler := handlers.NewParkingHandler(parkingService)
	vehicleHandler := handlers.NewVehicleHandler(vehicleService)
	mux := handlers.NewRouter(parkingHandler, vehicleHandler)

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
