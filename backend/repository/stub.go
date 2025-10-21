package repository

import "transportService/models"

// это заглушка, не нуждается в доработке
type ParkingStub struct {
	parkings map[int]string
}

func NewParkingStub() *ParkingStub {
	return &ParkingStub{make(map[int]string)}
}

func (s *ParkingStub) GetByID(id int) (models.Parking, error) {
	return models.Parking{}, nil
}

func (s *ParkingStub) GetAll() ([]models.Parking, error) {
	return []models.Parking{}, nil
}
