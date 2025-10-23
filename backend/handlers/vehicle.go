package handlers

import (
	"encoding/json"
	"net/http"
	"transportService/services"

	"go.uber.org/zap"
)

type VehicleHandler struct {
	service *services.VehicleService
}

func NewVehicleHandler(s *services.VehicleService) *VehicleHandler {
	return &VehicleHandler{service: s}
}

func (v *VehicleHandler) GetVehicles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	
	vehicles, err := v.service.GetVehicles()

	if err != nil {
		zap.L().Error("Error via getting vehicles", zap.Error(err))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	//TODO sending models in json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(vehicles); err != nil {
		zap.L().Error("Failed to encode vehicles", zap.Error(err))
	}
}
