package main

import (
	sqlc "TPE/db/sqlc"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func getturno(w http.ResponseWriter, r *http.Request, fecha pgtype.Date, hora int) {
	paramsGet := sqlc.GetTurnoParams{Fecha: fecha, Hora: int32(hora)}

	turno, err := queries.GetTurno(r.Context(), paramsGet)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(turno)
}

func deleteturno(w http.ResponseWriter, r *http.Request, fecha pgtype.Date, hora int) {
	turnoABorrar := sqlc.DeleteTurnoParams{Fecha: fecha, Hora: int32(hora)}

	err := queries.DeleteTurno(r.Context(), turnoABorrar)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func validarTurno(turno sqlc.Turno, esCreacion bool) error {
	trValido := turno.TratamientoID > 0
	clValido := turno.ClienteID > 0
	var horaValida bool
	if esCreacion {
		horaValida = turno.Hora < 20 && turno.Hora > 7
	} else {
		horaValida = true
	}

	if !trValido || !clValido || !horaValida {
		textoError := ""
		if !trValido {
			textoError += "El tratamiento ingresado no es válido.\n"
		}
		if !clValido {
			textoError += "El cliente ingresado no es válido.\n"
		}
		if !horaValida {
			textoError += "Solo se admiten turnos entre las 8hs y las 19hs.\n"
		}

		return fmt.Errorf(textoError)
	}
	return nil
}

func updateturno(w http.ResponseWriter, r *http.Request, fecha pgtype.Date, hora int) {
	var turnoToUpdate sqlc.Turno
	err := json.NewDecoder(r.Body).Decode(&turnoToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = validarTurno(turnoToUpdate, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updateParams := sqlc.UpdateTurnoParams{Fecha: fecha, Hora: int32(hora), TratamientoID: turnoToUpdate.TratamientoID, ClienteID: turnoToUpdate.ClienteID}
	UpdatedTurno, err := queries.UpdateTurno(r.Context(), updateParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(UpdatedTurno)
}

func createturno(w http.ResponseWriter, r *http.Request) {
	var nuevoturno sqlc.Turno
	err := json.NewDecoder(r.Body).Decode(&nuevoturno)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = validarTurno(nuevoturno, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	paramsCreacion := sqlc.CreateTurnoParams{Fecha: nuevoturno.Fecha, Hora: nuevoturno.Hora, TratamientoID: nuevoturno.TratamientoID, ClienteID: nuevoturno.ClienteID}
	nuevoturnoRow, err := queries.CreateTurno(r.Context(), paramsCreacion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(nuevoturnoRow)
	}
}

func getTurnos(w http.ResponseWriter, r *http.Request) {
	turnos, err := queries.GetTurnosCompleto(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(turnos)
}

func turnoHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	clave := parts[2]
	claves := strings.Split(clave, "&")
	fechaStr := claves[0]
	hora, err := strconv.Atoi(claves[1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	var fecha pgtype.Date
	fecha.Scan(fechaStr)

	switch r.Method {
	case http.MethodGet:
		getturno(w, r, fecha, hora)
	case http.MethodPut:
		updateturno(w, r, fecha, hora)
	case http.MethodDelete:
		deleteturno(w, r, fecha, hora)
	default:
		http.Error(w, "Método no permitido",
			http.StatusMethodNotAllowed)
	}
}

func turnosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createturno(w, r)
	case http.MethodGet:
		getTurnos(w, r)
	default:
		http.Error(w, "Método no permitido",
			http.StatusMethodNotAllowed)
	}
}
