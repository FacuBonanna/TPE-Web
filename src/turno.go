package main

import (
	sqlc "TPE/db/sqlc"
	"encoding/json"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func getturno(w http.ResponseWriter, r *http.Request) {
	var paramsGet sqlc.GetTurnoParams
	err := json.NewDecoder(r.Body).Decode(&paramsGet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	turno, err := queries.GetTurno(ctx, paramsGet)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(turno)
}

func deleteturno(w http.ResponseWriter, r *http.Request) {
	var turnoABorrar sqlc.DeleteTurnoParams
	err := json.NewDecoder(r.Body).Decode(&turnoABorrar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = queries.DeleteTurno(ctx, turnoABorrar)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func updateturno(w http.ResponseWriter, r *http.Request) {
	var turnoToUpdate sqlc.Turno
	err := json.NewDecoder(r.Body).Decode(&turnoToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateParams := sqlc.UpdateTurnoParams{Fecha: turnoToUpdate.Fecha, Hora: turnoToUpdate.Hora, TratamientoID: turnoToUpdate.TratamientoID, ClienteID: turnoToUpdate.ClienteID}

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updateParams)
}

func createturno(w http.ResponseWriter, r *http.Request) {
	var nuevoturno sqlc.Turno
	err := json.NewDecoder(r.Body).Decode(&nuevoturno)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	paramsCreacion := sqlc.CreateTurnoParams{Fecha: nuevoturno.Fecha, Hora: nuevoturno.Hora, TratamientoID: nuevoturno.TratamientoID, ClienteID: nuevoturno.ClienteID}
	nuevoturnoRow, err := queries.CreateTurno(ctx, paramsCreacion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(nuevoturnoRow)
	}
}

func turnoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createturno(w, r)
	case http.MethodGet:
		getturno(w, r)
	case http.MethodPut:
		updateturno(w, r)
	case http.MethodDelete:
		deleteturno(w, r)
	default:
		http.Error(w, "Método no permitido",
			http.StatusMethodNotAllowed)
	}
}
