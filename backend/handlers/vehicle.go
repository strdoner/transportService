package handlers

import (
	"go.uber.org/zap"
	"net/http"
	"transportService/services"
)

type VehicleHandler struct {
	service *services.VehicleService
}

func NewVehicleHandler(s *services.VehicleService) *VehicleHandler {
	return &VehicleHandler{service: s}
}

func (v *VehicleHandler) GetVehicles(w http.ResponseWriter, r *http.Request) {
	_, err := v.service.GetVehicles()

	if err != nil {
		zap.L().Error("Error via getting vehicles", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		//TODO sending reason
	}
	//TODO sending models in json
}
