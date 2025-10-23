package repository

import "transportService/models"

// это заглушка, не нуждается в доработке
type ParkingStub struct {
	parkings map[int]string
}

type VehicleStub struct{}

func NewParkingStub() *ParkingStub {
	return &ParkingStub{make(map[int]string)}
}

func NewVehicleStub() *VehicleStub {
	return &VehicleStub{}
}

func (s *ParkingStub) GetByID(id int) (models.Parking, error) {
	return models.Parking{}, nil
}

func (s *ParkingStub) GetAll() ([]models.Parking, error) {
	return []models.Parking{}, nil
}

func (v *VehicleStub) GetAll() ([]models.Vehicle, error) {
    return []models.Vehicle{}, nil
}

func (s *ParkingStub) Reserve(id int) error {
	panic("unimplemented")
}
