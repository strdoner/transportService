package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"transportService/services"

	"go.uber.org/zap"
)

type ParkingHandler struct {
	service *services.ParkingService
}

func NewParkingHandler(s *services.ParkingService) *ParkingHandler {
	return &ParkingHandler{service: s}
}

// Handles GET /parking/{id}  and POST /parking/{id}/reserve
func (h *ParkingHandler) HandleParking(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	if len(parts) < 2 || parts[0] != "parking" {
		http.NotFound(w, r)
		return
	}

	idStr := parts[1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid parking id", http.StatusBadRequest)
		return
	}

	// reserve
	if len(parts) == 3 && parts[2] == "reserve" {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		err := h.service.ReserveParking(id)
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
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "reserved"})
		return
	}

	// get parking by id
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	parking, err := h.service.GetParking(id)
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
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(parking); err != nil {
		zap.L().Error("Failed to encode parking", zap.Error(err))
	}
}

// GET  /parking
func (h *ParkingHandler) GetParkingLots(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	lots, err := h.service.GetParkingLots()
	if err != nil {
		zap.L().Error("Error via getting parking lots", zap.Error(err))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	//TODO sending models in json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(lots); err != nil {
		zap.L().Error("Failed to encode parking lots", zap.Error(err))
	}
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
