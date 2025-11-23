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
	rows, err := v.db.Query(`
		SELECT id, name, type, latitude, longitude, updated_at
		FROM vehicles
		ORDER BY id
	`)
	if err != nil {
		return nil, fmt.Errorf("GetAll query: %w", err)
	}
	defer rows.Close()

	var res []models.Vehicle
	for rows.Next() {
		var vehicle models.Vehicle
		var updated sql.NullTime
		if err := rows.Scan(&vehicle.Id, &vehicle.Name, &vehicle.Type, &vehicle.Latitude, &vehicle.Longitude, &updated); err != nil {
			return nil, fmt.Errorf("GetAll scan: %w", err)
		}
		if updated.Valid {
			vehicle.UpdatedAt = models.Timestamp{
				Seconds: updated.Time.Unix(),
				Nanos:   int32(updated.Time.Nanosecond()),
			}
		}
		res = append(res, vehicle)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAll rows: %w", err)
	}

	return res, nil
}
