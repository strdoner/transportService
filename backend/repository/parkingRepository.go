package repository

import (
	"database/sql"
	"fmt"
	"time"
	"transportService/models"
)

type IParkingRepository interface {
	GetAll() ([]models.Parking, error)
	GetByID(id int) (models.Parking, error)
}

type ParkingRepository struct {
	db *sql.DB
}

func NewParkingRepository(db *sql.DB) *ParkingRepository {
	return &ParkingRepository{db: db}
}

func (pr *ParkingRepository) GetByID(id int) (models.Parking, error) {
	var p models.Parking
	var created time.Time

	err := pr.db.QueryRow(`
		SELECT id, name, capacity, latitude, longitude, created_at
		FROM parking_lots
		WHERE id = $1
	`, id).Scan(&p.Id, &p.Name, &p.Capacity, &p.Latitude, &p.Longitude, &created)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, sql.ErrNoRows
		}
		return p, fmt.Errorf("GetByID query: %w", err)
	}

	p.CreatedAt = models.Timestamp{
		Seconds: created.Unix(),
		Nanos:   int32(created.Nanosecond()),
	}

	return p, nil
}

func (pr *ParkingRepository) GetAll() ([]models.Parking, error) {
	rows, err := pr.db.Query(`
		SELECT id, name, capacity, latitude, longitude, created_at
		FROM parking_lots
		ORDER BY id
	`)
	if err != nil {
		return nil, fmt.Errorf("GetAll query: %w", err)
	}
	defer rows.Close()

	var res []models.Parking
	for rows.Next() {
		var p models.Parking
		var created time.Time
		if err := rows.Scan(&p.Id, &p.Name, &p.Capacity, &p.Latitude, &p.Longitude, &created); err != nil {
			return nil, fmt.Errorf("GetAll scan: %w", err)
		}
		p.CreatedAt = models.Timestamp{
			Seconds: created.Unix(),
			Nanos:   int32(created.Nanosecond()),
		}
		res = append(res, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAll rows: %w", err)
	}
	return res, nil
}
