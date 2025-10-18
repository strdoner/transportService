package repository

type ParkingRepository interface {
	GetByID(id int) (interface{}, error)
	//other
}
