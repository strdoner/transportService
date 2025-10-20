#!/bin/sh
set -e

echo "Waiting for PostgreSQL at db:5432..."

while ! nc -z db 5432; do
  sleep 1
done

echo "PostgreSQL is ready!"

echo "Checking migration status..."

if [ -d "${MIGRATE_PATH:-/app/migrations}" ] && ls "${MIGRATE_PATH:-/app/migrations}"/*.sql >/dev/null 2>&1; then
  CURRENT_VERSION=$(migrate -path="${MIGRATE_PATH:-/app/migrations}" -database "${DATABASE_URL}" version 2>/dev/null || echo "error")

  if [ "$CURRENT_VERSION" = "error" ]; then
    echo "Could not get current migration version"
  else
    echo "Current migration version: $CURRENT_VERSION"
  fi

  echo "Running migrations..."
  if migrate -path="${MIGRATE_PATH:-/app/migrations}" -database "${DATABASE_URL}" up; then
    echo "Migrations completed successfully"
  else
    status=$?
    echo "Migration failed with code: $status"

    if migrate -path="${MIGRATE_PATH:-/app/migrations}" -database "${DATABASE_URL}" version 2>&1 | grep -q "dirty"; then
      echo "Database is in dirty state. Forcing version to latest..."
      LATEST_VERSION=$(ls "${MIGRATE_PATH:-/app/migrations}"/*.up.sql | wc -l)
      migrate -path="${MIGRATE_PATH:-/app/migrations}" -database "${DATABASE_URL}" force "$LATEST_VERSION"
      echo "Fixed dirty state. Version forced to: $LATEST_VERSION"
    else
      echo "Unknown migration error, exiting"
      exit 1
    fi
  fi
else
  echo "No migration files found, skipping migrations"
fi

echo "Starting application..."
exec /app/transport