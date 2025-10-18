package handlers

import (
	"go.uber.org/zap"
	"net/http"
	"transportService/middleware"
)

func NewRouter(parkingHandler *ParkingHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /parkings/{id}", parkingHandler.GetParking)
	mux.HandleFunc("GET /health", parkingHandler.CheckHealth)
	// TODO other handlers

	zap.L().Info("Routes registered successfully")

	loggedMux := middleware.LoggingMiddleware(mux)

	return loggedMux
}
