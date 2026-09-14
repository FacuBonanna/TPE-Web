APP_NAME := GLAM
CONTAINER_NAME := postgres-db

.PHONY: generate build clean test pre-test run-test post-test

all: 
	build

build: 
	@mkdir -p tmp
	@go build -o tmp/$(APP_NAME) ./src

clean:
	@rm -rf tmp

test: pre-test run-test post-test 

pre-test: build
	docker compose down -v 2>/dev/null || true
	docker compose up -d database
	@sleep 5
	docker compose up -d --force-recreate api 
	@echo "[PRE-TEST] Inyectando esquema SQL..."
	docker exec -i $(CONTAINER_NAME) psql -U postgres -d apirest < ./db/schema/schema.sql

run-test:
	@echo "[RUN-TEST] ejecutando pruebas"
	hurl --test ./requests.hurl

post-test:
	@echo "[POST-TEST] pruebas terminadas, limpiando"
	docker compose down -v
