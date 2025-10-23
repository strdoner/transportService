package handlers

import (
	"net/http"
	"transportService/middleware"

	"go.uber.org/zap"
)

func NewRouter(parkingHandler *ParkingHandler, vehicleHandler *VehicleHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/parking/", parkingHandler.HandleParking)
	mux.HandleFunc("/parking", parkingHandler.GetParkingLots)
	mux.HandleFunc("/health", parkingHandler.CheckHealth)
	mux.HandleFunc("/vehicles", vehicleHandler.GetVehicles)
	// TODO other handlers

	zap.L().Info("Routes registered successfully")

	loggedMux := middleware.LoggingMiddleware(mux)

	return loggedMux
}
