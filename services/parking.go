package services

import (
	"go.uber.org/zap"
	"transportService/repository"
)

type ParkingService struct {
	repo repository.ParkingRepository
}

func NewParkingService(repo repository.ParkingRepository) *ParkingService {
	service := &ParkingService{
		repo: repo,
	}

	zap.L().Info("Parking service created successfully")
	return service
}

func (s *ParkingService) GetParkingByID(id int) (interface{}, error) {

	return "not implemented yet", nil
}
