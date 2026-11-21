package main

import (
	sqlc "TPE/db/sqlc"
	views "TPE/views"
	"database/sql"
	"fmt"
	"net/http"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"

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
	http.HandleFunc("GET /clientes-filas", listarClientes)

	http.HandleFunc("POST /tratamiento", crearTratamiento)
	http.HandleFunc("DELETE /tratamiento/{id}", deleteTratamiento)
	http.HandleFunc("GET /tratamiento/{id}/edit", editarTratamientoFila)
	http.HandleFunc("GET /tratamiento/{id}", getTratamientoFila)
	http.HandleFunc("PUT /tratamiento/{id}", actualizarTratamiento)
	http.HandleFunc("POST /tratamientos", crearTratamiento)
	http.HandleFunc("DELETE /tratamientos/{id}", deleteTratamiento)
	http.HandleFunc("GET /tratamiento-filas", listarTratamientos)

	http.HandleFunc("POST /voucher", createVoucher)
	http.HandleFunc("DELETE /voucher/{id}", deleteVoucher)
	http.HandleFunc("POST /vouchers", createVoucher)
	http.HandleFunc("DELETE /vouchers/{id}", deleteVoucher)
	http.HandleFunc("GET /vouchers-filas", listarVouchers)
	http.HandleFunc("GET /voucher/{id}/edit", editarVoucherFila)
	http.HandleFunc("GET /voucher/{id}", getVoucherFila)
	http.HandleFunc("PUT /voucher/{id_voucher}", actualizarVoucher)

	http.HandleFunc("POST /turno", createTurno)
	http.HandleFunc("DELETE /turno/{fecha}/{hora}", deleteTurno)
	http.HandleFunc("DELETE /turno/{id}", deleteTurno)
	http.HandleFunc("POST /turnos", createTurno)
	http.HandleFunc("DELETE /turnos/{fecha}/{hora}", deleteTurno)
	http.HandleFunc("DELETE /turnos/{id}", deleteTurno)
	http.HandleFunc("GET /turno-filas", listarTurnos)
	http.HandleFunc("GET /turno/{fecha}/{hora}/edit", editarTurnoFila)
	http.HandleFunc("GET /turno/{fecha}/{hora}", getTurnoFila)
	http.HandleFunc("PUT /turno/{fecha}/{hora}", actualizarTurno)

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
func errorHTMX(w http.ResponseWriter, r *http.Request, texto string) {
    w.Header().Set("HX-Reswap", "none") // no toca la fila
    views.MensajeError(texto).Render(r.Context(), w)
}

func motivoBorrado(err error, entidad string) string {
    var pgErr *pgconn.PgError
    if errors.As(err, &pgErr) {
        switch pgErr.Code {
        case "23503": // foreign_key_violation
            return "No se puede eliminar: el " + entidad + " tiene registros asociados (" + pgErr.TableName + ")."
        case "23502": // not_null_violation
            return "No se puede eliminar: falta un dato obligatorio."
        }
    }
    return "No se pudo eliminar el " + entidad + " por un error inesperado."
}

func motivoCreacion(err error, entidad string) string {
    var pgErr *pgconn.PgError
    if errors.As(err, &pgErr) {
        switch pgErr.Code {
        case "23503": // foreign_key_violation
            return "El cliente o el tratamiento ingresado no existe."
        case "23505": // unique_violation
            return "Ya existe un " + entidad + " con esos datos."
        }
    }
    return "No se pudo crear el " + entidad + "."
}