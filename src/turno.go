package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	sqlc "TPE/db/sqlc"
	"TPE/views"

	"github.com/jackc/pgx/v5/pgtype"
)

func getTurnos(ctx context.Context) []sqlc.GetTurnosCompletoRow {
	tur, _ := queries.GetTurnosCompleto(ctx)
	return tur
}

func createTurno(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	fechaStr := r.FormValue("fecha")
	horaStr := r.FormValue("hora")
	idTratamientoStr := r.FormValue("tratamiento_id")
	idClienteStr := r.FormValue("cliente_documento")

	// Fecha: "2006-01-02" (formato de <input type="date">)
	t, err := time.Parse("2006-01-02", fechaStr)
	if err != nil {
		http.Error(w, "La fecha debe tener formato AAAA-MM-DD", http.StatusBadRequest)
		return
	}
	fecha := pgtype.Date{Time: t, Valid: true}

	hora64, err := strconv.ParseInt(horaStr, 10, 32)
	if err != nil {
		http.Error(w, "La hora debe ser un número entero válido", http.StatusBadRequest)
		return
	}

	idTratamiento, err := strconv.ParseInt(idTratamientoStr, 10, 64)
	if err != nil {
		http.Error(w, "El ID de tratamiento debe ser un número entero válido", http.StatusBadRequest)
		return
	}

	idCliente, err := strconv.ParseInt(idClienteStr, 10, 64)
	if err != nil {
		http.Error(w, "El ID de cliente debe ser un número entero válido", http.StatusBadRequest)
		return
	}

	_, err = queries.CreateTurno(r.Context(), sqlc.CreateTurnoParams{
		Fecha:            fecha,
		Hora:             int32(hora64),
		TratamientoID:    idTratamiento,
		ClienteDocumento: idCliente,
	})
	if err != nil {
		http.Error(w, "Error al crear el turno: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func getTurno(w http.ResponseWriter, r *http.Request) {
	fechaStr := r.FormValue("fecha")
	horaStr := r.FormValue("hora")

	// Fecha: "2006-01-02" (formato de <input type="date">)
	t, err := time.Parse("2006-01-02", fechaStr)
	if err != nil {
		http.Error(w, "La fecha debe tener formato AAAA-MM-DD", http.StatusBadRequest)
		return
	}
	fecha := pgtype.Date{Time: t, Valid: true}

	hora64, err := strconv.ParseInt(horaStr, 10, 32)
	if err != nil {
		http.Error(w, "La hora debe ser un número entero válido", http.StatusBadRequest)
		return
	}
	turno, err := queries.GetTurno(r.Context(), sqlc.GetTurnoParams{Fecha: fecha, Hora: int32(hora64)})

	clientes := getClientes(r.Context())
	tratamientos := getTratamientos(r.Context())
	turnos := sqlc.GetTurnosCompletoRow{Fecha: turno.Fecha, Hora: turno.Hora, ClienteDocumento: turno.ClienteDocumento, TratamientoID: turno.TratamientoID, NombreCliente: turno.NombreCliente, NombreTratamiento: turno.NombreTratamiento}
	turnosArr := []sqlc.GetTurnosCompletoRow{turnos}
	voucher := getVouchers(r.Context())

	err = views.Index(clientes, tratamientos, turnosArr, voucher).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error al renderizar la página", http.StatusInternalServerError)
	}
}
