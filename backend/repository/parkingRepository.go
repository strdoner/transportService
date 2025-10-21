package repository

import (
	"database/sql"
	"fmt"
	"transportService/models"
)

type IParkingRepository interface {
	GetByID(id int) (models.Parking, error)
	GetAll() ([]models.Parking, error)
}

type ParkingRepository struct {
	db *sql.DB
}

func NewParkingRepository(db *sql.DB) *ParkingRepository {
	return &ParkingRepository{db: db}
}

func (pr *ParkingRepository) GetById(id int) (models.Parking, error) {
	return models.Parking{}, fmt.Errorf("not implemented yet")
}

func (pr *ParkingRepository) GetAll() ([]models.Parking, error) {
	return []models.Parking{}, fmt.Errorf("not implemented yet")
}
