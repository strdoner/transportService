package repository

import (
	"database/sql"
	"time"
)

type IReservationRepository interface {
	GetCapacity(parkingID int) (int, error)
	CountActiveReservations(parkingID int, at time.Time) (int, error)
	CreateReservation(parkingID, vehicleID int, startsAt, expiresAt time.Time) error
}

type ReservationRepository struct {
	db *sql.DB
}

func NewReservationRepository(db *sql.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (r *ReservationRepository) GetCapacity(parkingID int) (int, error) {
	var capacity int
	err := r.db.QueryRow("SELECT capacity FROM parking_lots WHERE id=$1", parkingID).Scan(&capacity)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, sql.ErrNoRows
		}
		return 0, err
	}
	return capacity, nil
}

func (r *ReservationRepository) CountActiveReservations(parkingID int, at time.Time) (int, error) {
	var count int
	err := r.db.QueryRow(
		`SELECT COUNT(*) FROM reservations 
		 WHERE parking_id=$1 AND expires_at > $2`,
		parkingID, at,
	).Scan(&count)
	return count, err
}

func (r *ReservationRepository) CreateReservation(parkingID, vehicleID int, startsAt, expiresAt time.Time) error {
	_, err := r.db.Exec(
		`INSERT INTO reservations(parking_id, vehicle_id, starts_at, expires_at, created_at)
		 VALUES($1, $2, $3, $4, NOW())`,
		parkingID, vehicleID, startsAt, expiresAt,
	)
	return err
}
