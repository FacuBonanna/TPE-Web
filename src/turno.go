package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgconn"

	sqlc "TPE/db/sqlc"
	"TPE/views"
)

func getTurnos(ctx context.Context) []sqlc.GetTurnosCompletoRow {
	tur, _ := queries.GetTurnosCompleto(ctx)

	sort.Slice(tur, func(i, j int) bool {
		if !tur[i].Fecha.Equal(tur[j].Fecha) {
			return tur[i].Fecha.Before(tur[j].Fecha)
		}
		return tur[i].Hora < tur[j].Hora
	})
	return tur
}

func listarTurnos(w http.ResponseWriter, r *http.Request) {
	views.FilasTurno(getTurnos(r.Context())).Render(r.Context(), w)
}

func createTurno(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		errorHTMX(w, r, "No se pudo procesar el formulario.")
		return
	}

	fechaStr := strings.TrimSpace(r.FormValue("fecha"))
	horaStr := strings.TrimSpace(r.FormValue("hora"))
	idTratamientoStr := strings.TrimSpace(r.FormValue("tratamiento_id"))
	idClienteStr := strings.TrimSpace(r.FormValue("cliente_documento"))

	t, err1 := time.Parse("2006-01-02", fechaStr)
	if err1 != nil {
		errorHTMX(w, r, "La fecha ingresada no es válida.")
		return
	}
	hora64, err2 := strconv.ParseInt(horaStr, 10, 32)
	if err2 != nil {
		errorHTMX(w, r, "La hora ingresada no es válida.")
		return
	}
	idTratamiento, err3 := strconv.ParseInt(idTratamientoStr, 10, 64)
	if err3 != nil {
		errorHTMX(w, r, "El ID del tratamiento debe ser un número entero.")
		return
	}
	idCliente, err4 := strconv.ParseInt(idClienteStr, 10, 64)
	if err4 != nil {
		errorHTMX(w, r, "El documento del cliente debe ser un número entero.")
		return
	}
	fecha := t

	_, err := queries.CreateTurno(r.Context(), sqlc.CreateTurnoParams{
		Fecha:            fecha,
		Hora:             int32(hora64),
		TratamientoID:    idTratamiento,
		ClienteDocumento: idCliente,
	})
	if err != nil {
		log.Printf("createTurno: %v", err)
		errorHTMX(w, r, getMotivoTurno(err, "crear"))
		return
	}

	turnos := getTurnos(r.Context())
	if err := views.TablaTurnos(turnos).Render(r.Context(), w); err != nil {
		log.Printf("createTurno render: %v", err)
	}
}

func deleteTurno(w http.ResponseWriter, r *http.Request) {
	fechaStr := r.PathValue("fecha")
	horaStr := r.PathValue("hora")

	if fechaStr == "" && horaStr == "" {
		idStr := r.PathValue("id")
		parts := strings.Split(idStr, "_")
		if len(parts) == 2 {
			fechaStr = parts[0]
			horaStr = parts[1]
		}
	}

	t, err := time.Parse("2006-01-02", fechaStr)
	if err != nil {
		http.Error(w, "La fecha debe tener formato AAAA-MM-DD", http.StatusBadRequest)
		return
	}
	fecha := t

	hora64, err := strconv.ParseInt(horaStr, 10, 32)
	if err != nil {
		http.Error(w, "La hora debe ser un número entero válido", http.StatusBadRequest)
		return
	}

	err = queries.DeleteTurno(r.Context(), sqlc.DeleteTurnoParams{
		Fecha: fecha,
		Hora:  int32(hora64),
	})
	if err != nil {
		http.Error(w, "Error al eliminar el turno: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getTurno(w http.ResponseWriter, r *http.Request) {
	fechaStr := r.FormValue("fecha")
	horaStr := r.FormValue("hora")

	t, err := time.Parse("2006-01-02", fechaStr)
	if err != nil {
		views.FilasTurno(nil).Render(r.Context(), w)
		return
	}
	fecha := t

	hora64, err := strconv.ParseInt(horaStr, 10, 32)
	if err != nil {
		views.FilasTurno(nil).Render(r.Context(), w)
		return
	}
	turno, err := queries.GetTurno(r.Context(), sqlc.GetTurnoParams{Fecha: fecha, Hora: int32(hora64)})

	turnos := sqlc.GetTurnosCompletoRow{Fecha: turno.Fecha, Hora: turno.Hora, ClienteDocumento: turno.ClienteDocumento, TratamientoID: turno.TratamientoID, NombreCliente: turno.NombreCliente, NombreTratamiento: turno.NombreTratamiento}
	turnosArr := []sqlc.GetTurnosCompletoRow{turnos}

	err = views.FilasTurno(turnosArr).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error al renderizar la página", http.StatusInternalServerError)
	}
}

func editarTurnoFila(w http.ResponseWriter, r *http.Request) {
	fechaStr := r.PathValue("fecha")
	horaStr := r.PathValue("hora")

	t, err := time.Parse("2006-01-02", fechaStr)
	if err != nil {
		views.FilasTurno(nil).Render(r.Context(), w)
		return
	}
	fecha := t

	hora64, err := strconv.ParseInt(horaStr, 10, 32)
	if err != nil {
		views.FilasTurno(nil).Render(r.Context(), w)
		return
	}
	turno, err := queries.GetTurno(r.Context(), sqlc.GetTurnoParams{Fecha: fecha, Hora: int32(hora64)})
	if err != nil {
		errorHTMX(w, r, "El turno no existe.")
		return
	}

	_ = views.FilaTurnoEdit(turno).Render(r.Context(), w)
}

func getTurnoFila(w http.ResponseWriter, r *http.Request) {
	fechaStr := r.PathValue("fecha")
	horaStr := r.PathValue("hora")

	t, err := time.Parse("2006-01-02", fechaStr)
	if err != nil {
		views.FilasTurno(nil).Render(r.Context(), w)
		return
	}
	fecha := t

	hora64, err := strconv.ParseInt(horaStr, 10, 32)
	if err != nil {
		views.FilasTurno(nil).Render(r.Context(), w)
		return
	}
	tu, err := queries.GetTurno(r.Context(), sqlc.GetTurnoParams{Fecha: fecha, Hora: int32(hora64)})
	if err != nil {
		errorHTMX(w, r, "El turno no existe.")
		return
	}
	turno := sqlc.GetTurnosCompletoRow{Fecha: tu.Fecha, Hora: tu.Hora, ClienteDocumento: tu.ClienteDocumento, TratamientoID: tu.TratamientoID, NombreCliente: tu.NombreCliente, NombreTratamiento: tu.NombreTratamiento}
	_ = views.FilaTurno(turno).Render(r.Context(), w)
}

func actualizarTurno(w http.ResponseWriter, r *http.Request) {
	fecha, err := time.Parse("2006-01-02", r.PathValue("fecha"))
	if err != nil {
		errorHTMX(w, r, "Fecha inválida.")
		return
	}

	hora, err := strconv.ParseInt(r.PathValue("hora"), 10, 32)
	if err != nil {
		errorHTMX(w, r, "Hora inválida.")
		return
	}

	tratamientoId, err := strconv.ParseInt(r.FormValue("tratamiento_id"), 10, 64)
	if err != nil {
		errorHTMX(w, r, "Alguno de los valores ingresados no es válido.")
		return
	}

	clienteDoc, err := strconv.ParseInt(r.FormValue("cliente_documento"), 10, 64)
	if err != nil {
		errorHTMX(w, r, "Alguno de los valores ingresados no es válido.")
		return
	}

	t, err := queries.UpdateTurno(r.Context(), sqlc.UpdateTurnoParams{
		Fecha:            fecha,
		Hora:             int32(hora),
		TratamientoID:    tratamientoId,
		ClienteDocumento: clienteDoc,
	})
	if err != nil {
		log.Printf("actualizarTurno: %v", err)
		errorHTMX(w, r, motivoCreacion(err, "turno"))
		return
	}

	turno := sqlc.GetTurnosCompletoRow{
		Fecha:             t.Fecha,
		Hora:              t.Hora,
		ClienteDocumento:  t.ClienteDocumento,
		TratamientoID:     t.TratamientoID,
		NombreCliente:     t.NombreCliente,
		NombreTratamiento: t.NombreTratamiento,
	}
	_ = views.FilaTurno(turno).Render(r.Context(), w)
}

func getMotivoTurno(err error, funcion string) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23503":
			{
				switch pgErr.ConstraintName {
				case "turno_cliente":
					return "El cliente ingresado no existe."
				case "turno_tratamiento":
					return "El tratamiento ingresado no existe."
				}
			}
			return "El cliente o el tratamiento ingresado no existe." + pgErr.ConstraintName
		case "23505":
			return "Ya existe un turno en esa fecha."
		}
	}
	return "No se pudo " + funcion + " el turno."
}
