package repository

type ParkingStub struct {
	parkings map[int]string
}

func NewParkingStub() *ParkingStub {
	return &ParkingStub{make(map[int]string)}
}

func (s *ParkingStub) GetByID(id int) (interface{}, error) {
	return "not implemented yet", nil
}
