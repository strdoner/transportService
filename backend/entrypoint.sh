#!/usr/bin/env bash
set -euo pipefail

: "${DATABASE_URL:?DATABASE_URL must be set, e.g. postgres://user:pass@db:5432/dbname?sslmode=disable}"
: "${MIGRATE_PATH:=/migrations}"
: "${MAX_RETRIES:=30}"
: "${SLEEP_SECONDS:=1}"

# wait for postgres to be ready
retries=0
until pg_isready -d "$DATABASE_URL" >/dev/null 2>&1; do
  retries=$((retries+1))
  if [ "$retries" -ge "$MAX_RETRIES" ]; then
    echo "Postgres did not become ready in time"
    exit 1
  fi
  echo "Waiting for Postgres... ($retries/$MAX_RETRIES)"
  sleep "$SLEEP_SECONDS"
done

echo "Postgres is ready. Running migrations from ${MIGRATE_PATH}..."

# run migrations if there are any files in /migrations
if [ -d "${MIGRATE_PATH}" ] && [ "$(ls -A ${MIGRATE_PATH} 2>/dev/null | wc -l)" -gt 0 ]; then
  migrate -path="${MIGRATE_PATH}" -database "${DATABASE_URL}" up || {
    # migrate exits non-zero on "no change" in some versions; handle known states
    status=$?
    if [ $status -ne 0 ]; then
      echo "migrate returned code $status"
      # если ошибка не "no change", прерываем
      # иногда migrate возвращает 1 при idempotent; можно расширить обработку при необходимости
      # но по умолчанию завершаем с ошибкой чтобы не скрыть реальные проблемы
      exit $status
    fi
  }
else
  echo "No migrations found in ${MIGRATE_PATH}, skipping migrate."
fi

echo "Migrations complete. Starting application."

# Запуск приложения (передаём параметры в STDIN)
exec /app/transport