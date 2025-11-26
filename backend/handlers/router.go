package handlers

import (
	"net/http"
	"transportService/middleware"

	"go.uber.org/zap"
)

func NewRouter(parkingHandler *ParkingHandler, vehicleHandler *VehicleHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /parking/{id}", parkingHandler.GetParking)
	mux.HandleFunc("GET /parking", parkingHandler.GetParkingLots)
	mux.HandleFunc("GET /health", parkingHandler.CheckHealth)
	mux.HandleFunc("GET /vehicles", vehicleHandler.GetVehicles)
	mux.HandleFunc("POST /parking/{id}/reserve", parkingHandler.ReserveParking)

	zap.L().Info("Routes registered successfully")

	corsMux := middleware.CORSMiddleware(mux)
	loggedMux := middleware.LoggingMiddleware(corsMux)

	return loggedMux
}
