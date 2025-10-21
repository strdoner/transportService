package repository

import (
	"database/sql"
	"fmt"
	"transportService/models"
)

type IVehicleRepository interface {
	GetAll() ([]models.Vehicle, error)
}

type VehicleRepository struct {
	db *sql.DB
}

func NewVehicleRepository(db *sql.DB) *VehicleRepository {
	return &VehicleRepository{db: db}
}

func (v *VehicleRepository) GetAll() ([]models.Vehicle, error) {
	return []models.Vehicle{}, fmt.Errorf("not implemented yet")
}
