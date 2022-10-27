#!/bin/bash -e

APP_ENV=${APP_ENV:-local}

echo "[`date`] Running entrypoint script in the '${APP_ENV}' environment..."

echo "[`date`] Running DB migrations..."
migrate -database "postgres://postgres:postgres@10.5.0.5:5432/api_biblioteca?sslmode=disable" -path ./migrations up

echo "[`date`] Starting server..."
go run ./main.go