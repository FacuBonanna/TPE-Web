APP_NAME := GLAM
CONTAINER_NAME := postgres-db

.PHONY: generate build clean test pre-test run-test post-test

all: 
	build

test: pre-test run-test hurl-test post-test 

pre-test:
	docker compose down -v 2>/dev/null || true
	docker compose build api
	docker compose up -d database
	@sleep 5
	docker compose up -d --force-recreate api 
	@echo "[PRE-TEST] Inyectando esquema SQL"
	docker exec -i $(CONTAINER_NAME) psql -U postgres -d apirest < ./db/schema/schema.sql

run-test:
	@echo "[RUN-TEST] ejecutando pruebas"
	go test -v ./src

hurl-test:
	@echo "[HURL-TEST] ejecutando pruebas hurl"
	docker exec -i go-api hurl --test < ./requests.hurl

post-test:
	@echo "[POST-TEST] pruebas terminadas, limpiando"
	docker compose down -v
