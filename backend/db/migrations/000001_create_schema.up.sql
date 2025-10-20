CREATE TABLE parking_lots (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  capacity INTEGER NOT NULL CHECK (capacity > 0),
  latitude DOUBLE PRECISION NOT NULL,
  longitude DOUBLE PRECISION NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TYPE vehicle_type AS ENUM ('car', 'truck', 'motorcycle', 'bike', 'scooter');

CREATE TABLE vehicles (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  type vehicle_type NOT NULL,
  latitude DOUBLE PRECISION NOT NULL,
  longitude DOUBLE PRECISION NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE reservations (
  id SERIAL PRIMARY KEY,
  parking_id INTEGER NOT NULL REFERENCES parkings(id),
  vehicle_id INTEGER NOT NULL REFERENCES vehicles(id),
  starts_at TIMESTAMPTZ NOT NULL,
  expires_at TIMESTAMPTZ NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  CHECK (expires_at > starts_at)
);