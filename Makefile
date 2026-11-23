APP_NAME := GLAM
CONTAINER_NAME := postgres-db

.PHONY: generate build clean test pre-test run-test post-test inicializacion finalizar

all: 
	inicializacion

test_base: pre-test run-test post-test 

test_api: pre-test hurl-test post-test

generate: 
	sqlc generate
	templ generate

pre-test:
	docker compose down -v 2>/dev/null || true
	docker compose build api
	docker compose up -d database	
	docker compose up -d --wait database
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

rebuild:
	docker compose down
	docker compose build api
	docker compose up api

# genera, levanta y en el detach baja
inicializacion: generate
	@trap 'echo "[FINALIZAR] limpiando"; docker compose down -v' EXIT INT TERM; \
	docker compose down -v && \
	docker compose build api && \
	docker compose up -d --wait database && \
	docker exec -i $(CONTAINER_NAME) psql -U postgres -d apirest < ./db/schema/schema.sql && \
	docker compose up api

finalizar:
	docker compose down -v

levantar:
	docker compose down -v 
	docker compose build api
	docker compose up -d --wait database 
	docker exec -i $(CONTAINER_NAME) psql -U postgres -d apirest < ./db/schema/schema.sql 
	docker compose up api
