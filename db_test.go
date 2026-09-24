package main

import (
	"testing"
	sqlc "TPE/db/sqlc"
	"context"
	"fmt"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"

)

func abrirDB() (*sql.DB, error) {

	dbString := fmt.Sprintf("host=postgres-db port=5432 user=postgres password=postgres dbname=apirest sslmode=disable")
	db, err := sql.Open("pgx", dbString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil { // Verifica que la BD responde
		return nil, err
	}
	db.SetMaxOpenConns(25) // Configura el tamaño del pool
	return db, nil
}

//cliente 
func TestQueries_CRUD(t *testing.T) {
	/*db, err := abrirDB()
	if err != nil {
		fmt.Printf("No se pudo conectar a la DB\n")
		fmt.Print(err.Error())
		return
	}*/
	t.Run ("Crear cliente", func(t *testing.T) {  
	paramsCreacion := sqlc.CreateUserParams{Nombre: "valentina", Apellido: "bisogni", Deuda: 10000, NroTelefono: 123456}
	
	_, err := queries.CreateUser(ctx, paramsCreacion)
	if err != nil {
		t.Errorf("No se creo el cliente")
	} })
}