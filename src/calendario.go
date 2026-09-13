package main

import (
	sqlc "TPE/db/sqlc"
	"encoding/json"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func getCalendario(w http.ResponseWriter, r *http.Request) {
	var paramsGet sqlc.GetCalendarParams
	err := json.NewDecoder(r.Body).Decode(&paramsGet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	calendario, err := queries.GetCalendar(ctx, paramsGet)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(calendario)
}

func deleteCalendario(w http.ResponseWriter, r *http.Request) {
	var calendarioABorrar sqlc.DeleteCalendarParams
	err := json.NewDecoder(r.Body).Decode(&calendarioABorrar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = queries.DeleteCalendar(ctx, calendarioABorrar)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func updateCalendario(w http.ResponseWriter, r *http.Request) {
	var calendarioToUpdate sqlc.Calendario
	err := json.NewDecoder(r.Body).Decode(&calendarioToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateParams := sqlc.UpdateCalendarParams{Fecha: calendarioToUpdate.Fecha, Hora: calendarioToUpdate.Hora, TratamientoID: calendarioToUpdate.TratamientoID, ClienteID: calendarioToUpdate.ClienteID}
	err = queries.UpdateCalendar(ctx, updateParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}

func createCalendario(w http.ResponseWriter, r *http.Request) {
	var nuevoCalendario sqlc.Calendario
	err := json.NewDecoder(r.Body).Decode(&nuevoCalendario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	paramsCreacion := sqlc.CreateCalendarParams{Fecha: nuevoCalendario.Fecha, Hora: nuevoCalendario.Hora, TratamientoID: nuevoCalendario.TratamientoID, ClienteID: nuevoCalendario.ClienteID}
	nuevoCalendarioRow, err := queries.CreateCalendar(ctx, paramsCreacion)
	json.NewEncoder(w).Encode(nuevoCalendarioRow)
}

func calendarioHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createCalendario(w, r)
	case http.MethodGet:
		getCalendario(w, r)
	case http.MethodPut:
		updateCalendario(w, r)
	case http.MethodDelete:
		deleteCalendario(w, r)
	default:
		http.Error(w, "Método no permitido",
			http.StatusMethodNotAllowed)
	}
}
