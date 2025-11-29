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
	vehicles, err := v.service.GetVehicles() // TODO add field in db for number of car etc. A123BC
	if err != nil {
		zap.L().Error("Error via getting vehicles", zap.Error(err))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(vehicles); err != nil {
		zap.L().Error("Failed to encode vehicles", zap.Error(err))
	}
}
