package main

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
	"log"
	"sort"

	sqlc "TPE/db/sqlc"
	"TPE/views"

	"github.com/jackc/pgx/v5/pgtype"
)

func getTurnos(ctx context.Context) []sqlc.GetTurnosCompletoRow {
tur, _ := queries.GetTurnosCompleto(ctx)

	sort.Slice(tur, func(i, j int) bool {
	if !tur[i].Fecha.Time.Equal(tur[j].Fecha.Time) {
		return tur[i].Fecha.Time.Before(tur[j].Fecha.Time)
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


    // validar valores
    t, err1 := time.Parse("2006-01-02", fechaStr)
    hora64, err2 := strconv.ParseInt(horaStr, 10, 32)
    idTratamiento, err3 := strconv.ParseInt(idTratamientoStr, 10, 64)
    idCliente, err4 := strconv.ParseInt(idClienteStr, 10, 64)
    if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
        errorHTMX(w, r, "Alguno de los valores ingresados no es válido.")
        return
    }
    fecha := pgtype.Date{Time: t, Valid: true}

    _, err := queries.CreateTurno(r.Context(), sqlc.CreateTurnoParams{
        Fecha:            fecha,
        Hora:             int32(hora64),
        TratamientoID:    idTratamiento,
        ClienteDocumento: idCliente,
    })
    if err != nil {
        log.Printf("createTurno: %v", err)
        errorHTMX(w, r, motivoCreacion(err, "turno"))
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
	fecha := pgtype.Date{Time: t, Valid: true}

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

	// Fecha: "2006-01-02" (formato de <input type="date">)
	t, err := time.Parse("2006-01-02", fechaStr)
	if err != nil {
		views.FilasTurno(nil).Render(r.Context(), w)
		return
	}
	fecha := pgtype.Date{Time: t, Valid: true}

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

	// Fecha: "2006-01-02" (formato de <input type="date">)
	t, err := time.Parse("2006-01-02", fechaStr)
	if err != nil {
		views.FilasTurno(nil).Render(r.Context(), w)
		return
	}
	fecha := pgtype.Date{Time: t, Valid: true}

	hora64, err := strconv.ParseInt(horaStr, 10, 32)
	if err != nil {
		views.FilasTurno(nil).Render(r.Context(), w)
		return
	}
	turno, err := queries.GetTurno(r.Context(), sqlc.GetTurnoParams{Fecha: fecha, Hora: int32(hora64)})

	_ = views.FilaTurnoEdit(turno).Render(r.Context(), w)
}

func getTurnoFila(w http.ResponseWriter, r *http.Request) {
	fechaStr := r.PathValue("fecha")
	horaStr := r.PathValue("hora")

	// Fecha: "2006-01-02" (formato de <input type="date">)
	t, err := time.Parse("2006-01-02", fechaStr)
	if err != nil {
		views.FilasTurno(nil).Render(r.Context(), w)
		return
	}
	fecha := pgtype.Date{Time: t, Valid: true}

	hora64, err := strconv.ParseInt(horaStr, 10, 32)
	if err != nil {
		views.FilasTurno(nil).Render(r.Context(), w)
		return
	}
	tu, err := queries.GetTurno(r.Context(), sqlc.GetTurnoParams{Fecha: fecha, Hora: int32(hora64)})
	turno := sqlc.GetTurnosCompletoRow{Fecha: tu.Fecha, Hora: tu.Hora, ClienteDocumento: tu.ClienteDocumento, TratamientoID: tu.TratamientoID, NombreCliente: tu.NombreCliente, NombreTratamiento: tu.NombreTratamiento}
	_ = views.FilaTurno(turno).Render(r.Context(), w)
}

func actualizarTurno(w http.ResponseWriter, r *http.Request) {
	fechaStr := r.PathValue("fecha")
	horaStr := r.PathValue("hora")

	time, err := time.Parse("2006-01-02", fechaStr)
	fecha := pgtype.Date{Time: time, Valid: true}

	hora, err := strconv.ParseInt(horaStr, 10, 64)
	if err != nil {
		http.Error(w, "Hora inválida", http.StatusBadRequest)
		return
	}

	tratamientoIdStr := r.FormValue("tratamiento_id")
	tratamientoId, err := strconv.ParseInt(tratamientoIdStr, 10, 64)
	if err != nil {
		http.Error(w, "ID del tratamiento inválido", http.StatusBadRequest)
		return
	}

	clienteDocStr := r.FormValue("cliente_documento")
	clienteDoc, err := strconv.ParseInt(clienteDocStr, 10, 64)
	if err != nil {
		http.Error(w, "Documento del cliente inválido", http.StatusBadRequest)
		return
	}
	t, err := queries.UpdateTurno(r.Context(), sqlc.UpdateTurnoParams{
		Fecha:            fecha,
		Hora:             int32(hora),
		TratamientoID:    tratamientoId,
		ClienteDocumento: clienteDoc,
	})
	if err != nil {
		http.Error(w, "Error al actualizar el turno: "+err.Error(), http.StatusInternalServerError)
		return
	}
	turno := sqlc.GetTurnosCompletoRow{Fecha: t.Fecha, Hora: t.Hora, ClienteDocumento: t.ClienteDocumento, TratamientoID: t.ClienteDocumento, NombreCliente: t.NombreCliente, NombreTratamiento: t.NombreTratamiento}
	_ = views.FilaTurno(turno).Render(r.Context(), w)
}
