CREATE TABLE parking (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  capacity INT NOT NULL,
  available_spots INT NOT NULL,
  latitude DOUBLE PRECISION,
  longitude DOUBLE PRECISION,
  created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE vehicles (
  id SERIAL PRIMARY KEY,
  plate TEXT,
  type TEXT,
  latitude DOUBLE PRECISION,
  longitude DOUBLE PRECISION,
  status TEXT,
  updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE reservations (
  id SERIAL PRIMARY KEY,
  parking_id INT REFERENCES parking(id) ON DELETE CASCADE,
  vehicle_id INT REFERENCES vehicles(id) ON DELETE SET NULL,
  status TEXT NOT NULL,
  reserved_at TIMESTAMPTZ DEFAULT now()
);