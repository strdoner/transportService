package services

import (
	"go.uber.org/zap"
	"transportService/models"
	"transportService/repository"
)

type ParkingService struct {
	repo repository.IParkingRepository
}

func NewParkingService(repo repository.IParkingRepository) *ParkingService {
	service := &ParkingService{
		repo: repo,
	}

	zap.L().Info("Parking service created successfully")
	return service
}

func (s *ParkingService) GetParkingLots() ([]models.Parking, error) {
	return s.repo.GetAll()
}
