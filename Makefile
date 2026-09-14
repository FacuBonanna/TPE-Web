APP_NAME := "GLAM"
BD_URL := 
CONTAINER_NAME := postgres-db

generate:
	@sqlc generate

build: generate
	@mkdir -p tmp
	@go build -o tmp/$(APP_NAME) .

# Limpia los artefactos de construcción
clean:
@rm -rf tmp

.PHONY generate build clean test pre-test run-test post-test

test: pre-test run-test post-test 

#tareas previas, levantar el docker (¿hay q borrar anteriores?), generar el sql, etc
pre-test: build
	docker compose down -v 2>/dev/null || true
	docker compose up -d --force-recreate api

	@echo "[PRE-TEST] Inyectando esquema SQL..."
	docker exec -i $(CONTAINER_NAME) psql -U postgres -d apirest < db/schema.sql

#se ejecuta con make test

run-test:
	@echo "[RUN-TEST] ejecutando pruebas"
	hurl --test ./requests.hurl




#dar de baja el docker 
post-test:
	@echo "[POST-TEST] pruebas terminadas, limpiando"
	docker compose down -v


