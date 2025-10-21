package repository

import (
	"database/sql"
	"fmt"
)

type IReservationRepository interface {
	ReserveParking(id int) error
}

type ReservationRepository struct {
	db *sql.DB
}

func NewReservationRepository(db *sql.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (r *ReservationRepository) ReserveParking(id int) error {
	return fmt.Errorf("not implemented yet")
}
