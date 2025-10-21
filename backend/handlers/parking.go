package handlers

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"transportService/services"
)

type ParkingHandler struct {
	service *services.ParkingService
}

func NewParkingHandler(s *services.ParkingService) *ParkingHandler {
	return &ParkingHandler{service: s}
}

func (h *ParkingHandler) GetParking(w http.ResponseWriter, r *http.Request) {
	//TODO API LOGIC
}

func (h *ParkingHandler) GetParkingLots(w http.ResponseWriter, r *http.Request) {
	_, err := h.service.GetParkingLots()
	if err != nil {
		zap.L().Error("Error via getting parking lots", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		//TODO sending reason
	}
	//TODO sending models in json
}

func (h *ParkingHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":  "ok",
		"service": "parking-service",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		zap.L().Error("Failed to encode health response", zap.Error(err))
	}
}
