package services

import (
	"fmt"
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
	veh, err := s.repo.GetAll()
    if err != nil {
        return nil, fmt.Errorf("failed to fetch vehicle: %w", err)
    }

    if len(veh) == 0 {
        return []models.Vehicle{}, nil
    }

    return veh, nil
}
