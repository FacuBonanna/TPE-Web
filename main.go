package main

import (
	sqlc "TPE/db/sqlc"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type repository struct {
	db *sql.DB
}

func abrirDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=server password=admin dbname=baseprueba sslmode=disable") //CAMBIAR USER,PASS, ETC.
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil { // Verifica que la BD responde
		return nil, err
	}
	db.SetMaxOpenConns(25) // Configura el tamaño del pool
	return db, nil
}

func getVoucher(w http.ResponseWriter, r *http.Request, id int) {

	voucher, err := queries.GetVoucher(ctx, int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Print(voucher)
}

func deleteVoucher(w http.ResponseWriter, r *http.Request, id int) {

	err := queries.DeleteVoucher(ctx, int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

}

func createVoucher(w http.ResponseWriter, r *http.Request) {
	var nuevoVoucher sqlc.Voucher
	err := json.NewDecoder(r.Body).Decode(&nuevoVoucher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	paramsCreacion := sqlc.CreateVoucherParams{RegaladorID: nuevoVoucher.ClienteID, TratamientoID: nuevoVoucher.RegaladorID, ClienteID: nuevoVoucher.TratamientoID}
	queries.CreateVoucher(ctx, paramsCreacion)
}

func updateVoucher(w http.ResponseWriter, r *http.Request, id int) {
	var updatedVoucher sqlc.Voucher
	err := json.NewDecoder(r.Body).Decode(&updatedVoucher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = queries.DeleteVoucher(ctx, int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

}

func voucherHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[2])

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		getVoucher(w, r, id)
	case http.MethodPut:
		updateVoucher(w, r, id)
	case http.MethodDelete:
		deleteVoucher(w, r, id)
	default:
		http.Error(w, "Método no permitido",
			http.StatusMethodNotAllowed)
	}
}

func vouchersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createVoucher(w, r)
	default:
		http.Error(w, "Method not allowed",
			http.StatusMethodNotAllowed)
	}
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

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}
