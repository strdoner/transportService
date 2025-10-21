package services

import (
	"transportService/models"
	"transportService/repository"
)

type VehicleService struct {
	repo repository.IVehicleRepository
}

func NewVehicleService(repo repository.IVehicleRepository) *VehicleService {
	service := &VehicleService{
		repo: repo,
	}

	return service
}

func (s *VehicleService) GetVehicles() ([]models.Vehicle, error) {
	return s.repo.GetAll()
}
