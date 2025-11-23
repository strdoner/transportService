package services

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"transportService/models"
	"transportService/repository"

	"go.uber.org/zap"
)

type ParkingService struct {
	parkingRepo     repository.IParkingRepository
    reservationRepo repository.IReservationRepository
}

func NewParkingService(parkingRepo repository.IParkingRepository, reservationRepo repository.IReservationRepository) *ParkingService {
    service := &ParkingService{
        parkingRepo:     parkingRepo,
        reservationRepo: reservationRepo,
    }

	zap.L().Info("Parking service created successfully")
	return service
}

func (s *ParkingService) GetParkingLots() ([]models.Parking, error) {
    lots, err := s.parkingRepo.GetAll()
    if err != nil {
        return nil, fmt.Errorf("failed to fetch parking lots: %w", err)
    }

    if len(lots) == 0 {
        return []models.Parking{}, nil
    }

    return lots, nil
}

func (s *ParkingService) GetParking(id int) (models.Parking, error) {
	p, err := s.parkingRepo.GetByID(id)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return p, fmt.Errorf("parking with id %d not found", id)
        }
        return p, fmt.Errorf("failed to get parking: %w", err)
    }
	return p, nil
}

func (s *ParkingService) ReserveParking(parkingID, vehicleID int, startsAt, expiresAt time.Time) error {
	capacity, err := s.reservationRepo.GetCapacity(parkingID)
	if err != nil {
		return err
	}

	reserved, err := s.reservationRepo.CountActiveReservations(parkingID, time.Now())
	if err != nil {
		return err
	}

	if reserved >= capacity {
		return errors.New("no available spots")
	}

	return s.reservationRepo.CreateReservation(parkingID, vehicleID, startsAt, expiresAt)
}