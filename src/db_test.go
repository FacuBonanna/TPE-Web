package main

import (
sqlc "TPE/db/sqlc"
	"context"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"


)


//cliente 
func TestQueriesCliente_CRUD(t *testing.T) {
		// 1. FORZAMOS la conexión local solo para el entorno de pruebas.
	// Esto no afecta a tu main.go en producción porque esta función solo corre en 'go test'.
	dbString := "host=localhost port=5432 user=postgres password=postgres dbname=apirest sslmode=disable"
	db, err := sql.Open("pgx", dbString)
	if err != nil {
		t.Fatalf("No se pudo abrir la DB local para el test: %v", err)
	}
	if err := db.Ping(); err != nil {
		t.Fatalf("La DB de Docker no responde en localhost:5432. ¿Está el puerto expuesto? Error: %v", err)
	}
	defer db.Close()

	// 2. Inicializamos las variables globales que declaraste en tu main.go
	queries = sqlc.New(db)
	ctx = context.Background()

	t.Run("test", x)
	t.Run ("Crear cliente", func(t *testing.T) {  
	paramsCreacion := sqlc.CreateUserParams{Nombre: "valentina", Apellido: "bisogni", Deuda: 10000, NroTelefono: 123456}
	
	_, err := queries.CreateUser(ctx, paramsCreacion)
	if err != nil {
		t.Errorf("No se creo el cliente")
	} else { fmt.Print("se creo exitosamente el cliente")}})
}

func x(t *testing.T){}
