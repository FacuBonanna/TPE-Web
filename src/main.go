package main

import (
	sqlc "TPE/db/sqlc"
	"context"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type repository struct {
	db *sql.DB
}

func abrirDB() (*sql.DB, error) {

	dbString := fmt.Sprintf("host=postgres-db port=5432 user=postgres password=postgres dbname=apirest sslmode=disable")
	db, err := sql.Open("pgx", dbString)
	fmt.Printf(dbString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil { // Verifica que la BD responde
		return nil, err
	}
	db.SetMaxOpenConns(25) // Configura el tamaño del pool
	return db, nil
}

var queries *sqlc.Queries
var ctx context.Context

func main() {
	db, err := abrirDB()
	if err != nil {
		fmt.Printf("No se pudo conectar a la DB\n")
		fmt.Printf(err.Error())
		return
	}
	queries = sqlc.New(db)
	ctx = context.Background()

	fileServer := http.FileServer(http.Dir("./"))
	http.Handle("/", fileServer)
	http.HandleFunc("/cliente/", clienteHandler)
	http.HandleFunc("/cliente", clientesHandler)
	http.HandleFunc("/voucher/", voucherHandler)
	http.HandleFunc("/voucher", vouchersHandler)
	http.HandleFunc("/tratamiento/", tratamientoHandler)
	http.HandleFunc("/tratamiento", tratamientosHandler)
	http.HandleFunc("/turno", turnoHandler)

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}
