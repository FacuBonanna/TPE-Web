package main

import (
	sqlc "TPE/db/sqlc"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type repository struct {
	db *sql.DB
}

func abrirDB() (*sql.DB, error) {
	_ = godotenv.Load()
	puerto := os.Getenv("puerto")
	usuario := os.Getenv("usuario")
	dbname := os.Getenv("dbname")
	password := os.Getenv("password")
	host := os.Getenv("host")

	dbString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, puerto, usuario, password, dbname)
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
	http.HandleFunc("/turno/", turnoHandler)

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}
