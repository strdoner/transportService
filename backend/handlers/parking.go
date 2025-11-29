package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
	"transportService/services"

	"go.uber.org/zap"
)

type ParkingHandler struct {
	service *services.ParkingService
}

func NewParkingHandler(s *services.ParkingService) *ParkingHandler {
	return &ParkingHandler{service: s}
}

func (h *ParkingHandler) GetParking(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid parking id", http.StatusBadRequest)
		return
	}

	parking, err := h.service.GetParking(id) // TODO add field not in db for checking available capacity + query from db (ask for more info)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "parking not found", http.StatusNotFound)
			return
		}
		zap.L().Error("Error getting parking", zap.Int("id", id), zap.Error(err))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(parking); err != nil {
		zap.L().Error("failed to encode parking response", zap.Error(err))
	}
}

func (h *ParkingHandler) GetParkingLots(w http.ResponseWriter, r *http.Request) {
	lots, err := h.service.GetParkingLots()
	if err != nil {
		zap.L().Error("failed to get parking lots", zap.Error(err))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(lots); err != nil {
		zap.L().Error("failed to encode parking lots response", zap.Error(err))
	}
}

func (h *ParkingHandler) ReserveParking(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid parking id", http.StatusBadRequest)
		return
	}

	type ReserveRequest struct { //TODO Date of reserving + field in db
		VehicleID int       `json:"vehicle_id"`
		StartsAt  time.Time `json:"starts_at"`
		ExpiresAt time.Time `json:"expires_at"`
	}

	var req ReserveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = h.service.ReserveParking(id, req.VehicleID, req.StartsAt, req.ExpiresAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "parking not found", http.StatusNotFound)
			return
		}
		if strings.Contains(err.Error(), "no available") || strings.Contains(err.Error(), "full") {
			http.Error(w, "no available spots", http.StatusConflict)
			return
		}
		zap.L().Error("Error reserving parking", zap.Int("id", id), zap.Error(err))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "reserved"})
}

func (h *ParkingHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":  "ok",
		"service": "parking-service",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		zap.L().Error("Failed to encode health response", zap.Error(err))
	}
}
