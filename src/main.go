package main

import (
	sqlc "TPE/db/sqlc"
	views "TPE/views"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var queries *sqlc.Queries

func abrirDB() (*sql.DB, error) {
	dbString := "host=postgres-db port=5432 user=postgres password=postgres dbname=apirest sslmode=disable"

	db, err := sql.Open("pgx", dbString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)

	return db, nil
}

func main() {
	dbase, err := abrirDB()
	if err != nil {
		fmt.Println("No se pudo conectar a la DB:", err)
		return
	}
	defer dbase.Close()

	queries = sqlc.New(dbase)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", servidor)

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)

	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)

	}
}
func servidor(w http.ResponseWriter, r *http.Request) {

	//	turnos := make([]db.Turno, 0, 0)
	//	vouchers := make([]db.Voucher, 0, 0)
	//	clientes := make([]db.Cliente, 0, 0)
	clientes := getClientes(r.Context())

	err := views.Index(clientes).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error al renderizar la página", http.StatusInternalServerError)
		return
	}
}
