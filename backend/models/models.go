package models

type Timestamp struct {
	Seconds int64 `json:"seconds"`
	Nanos   int32 `json:"nanos"`
}

type Vehicle struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	UpdatedAt Timestamp `json:"updated_at"`
}

type Parking struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Capacity  int32     `json:"capacity"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CreatedAt Timestamp `json:"created_at"`
}

type Reservation struct {
	Id        int       `json:"id"`
	ParkingId int       `json:"parking_id"`
	VehicleId int       `json:"vehicle_id"`
	StartsAt  Timestamp `json:"starts_at"`
	ExpiresAt Timestamp `json:"expires_at"`
	CreatedAt Timestamp `json:"created_at"`
}
