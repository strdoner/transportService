package services

import (
	"transportService/models"
	"transportService/repository"

	"go.uber.org/zap"
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

func (s *ParkingService) GetParking(id int) (models.Parking, error) {
	return s.repo.GetByID(id)
}

func (s *ParkingService) ReserveParking(id int) error {
	return s.repo.Reserve(id)
}