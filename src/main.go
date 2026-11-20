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
	//http.HandleFunc("Get /{$}", servidor)
	http.HandleFunc("POST /cliente", crearCliente)
	http.HandleFunc("DELETE /cliente/{id}", deleteCliente)
	http.HandleFunc("GET /cliente/{id}/edit", editarClienteFila)
	http.HandleFunc("GET /cliente/{id}", getClienteFila)
	http.HandleFunc("PUT /cliente/{id}", actualizarCliente)
	http.HandleFunc("POST /clientes", crearCliente)
	http.HandleFunc("DELETE /clientes/{id}", deleteCliente)

	http.HandleFunc("POST /tratamiento", crearTratamiento)
	http.HandleFunc("DELETE /tratamiento/{id}", deleteTratamiento)
	http.HandleFunc("GET /tratamiento/{id}/edit", editarTratamientoFila)
	http.HandleFunc("GET /tratamiento/{id}", getTratamientoFila)
	http.HandleFunc("PUT /tratamiento/{id}", actualizarTratamiento)
	http.HandleFunc("POST /tratamientos", crearTratamiento)
	http.HandleFunc("DELETE /tratamientos/{id}", deleteTratamiento)

	http.HandleFunc("POST /voucher", createVoucher)
	http.HandleFunc("DELETE /voucher/{id}", deleteVoucher)
	http.HandleFunc("POST /vouchers", createVoucher)
	http.HandleFunc("DELETE /vouchers/{id}", deleteVoucher)

	http.HandleFunc("POST /turno", createTurno)
	http.HandleFunc("DELETE /turno/{fecha}/{hora}", deleteTurno)
	http.HandleFunc("DELETE /turno/{id}", deleteTurno)
	http.HandleFunc("POST /turnos", createTurno)
	http.HandleFunc("DELETE /turnos/{fecha}/{hora}", deleteTurno)
	http.HandleFunc("DELETE /turnos/{id}", deleteTurno)

	http.HandleFunc("GET /search-client", getCliente)
	http.HandleFunc("GET /search-tratamiento", getTratamiento)
	http.HandleFunc("GET /search-voucher", getVoucher)
	http.HandleFunc("GET /search-turno", getTurno)

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)

	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)

	}
}
func servidor(w http.ResponseWriter, r *http.Request) {
	clientes := getClientes(r.Context())
	tratamientos := getTratamientos(r.Context())
	vouchers := getVouchers(r.Context())
	turnos := getTurnos(r.Context())

	err := views.Index(clientes, tratamientos, turnos, vouchers).Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error al renderizar la página", http.StatusInternalServerError)
		return
	}

}
