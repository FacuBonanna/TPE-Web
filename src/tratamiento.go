package main

import (
	sqlc "TPE/db/sqlc"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func getTratamiento(w http.ResponseWriter, r *http.Request, id int) {

	tratamiento, err := queries.GetTreatment(ctx, int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(tratamiento)
}

func deleteTratamiento(w http.ResponseWriter, r *http.Request, id int) {

	err := queries.DeleteTreatment(ctx, int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func updateTratamiento(w http.ResponseWriter, r *http.Request, id int) {
	var tratamientoToUpdate sqlc.Tratamiento
	err := json.NewDecoder(r.Body).Decode(&tratamientoToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	paramsTratamientoToUpdate := sqlc.UpdateTreatmentParams{ID: int64(id), Nombre: tratamientoToUpdate.Nombre, DescripcionCorta: tratamientoToUpdate.DescripcionCorta, Costo: tratamientoToUpdate.Costo}
	updatedTreatment, err := queries.UpdateTreatment(ctx, paramsTratamientoToUpdate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedTreatment)
}

func createTratamiento(w http.ResponseWriter, r *http.Request) {
	var nuevoTratamiento sqlc.Tratamiento
	err := json.NewDecoder(r.Body).Decode(&nuevoTratamiento)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Printf(err.Error())
		return
	}
	paramsCreacion := sqlc.CreateTreatmentParams{Nombre: nuevoTratamiento.Nombre, DescripcionCorta: nuevoTratamiento.DescripcionCorta, Costo: nuevoTratamiento.Costo}
	nuevoTratamientoRow, err := queries.CreateTreatment(ctx, paramsCreacion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(nuevoTratamientoRow)
	}

}

func tratamientoHandler(w http.ResponseWriter, r *http.Request) {
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
		getTratamiento(w, r, id)
	case http.MethodPut:
		updateTratamiento(w, r, id)
	case http.MethodDelete:
		deleteTratamiento(w, r, id)
	default:
		http.Error(w, "Método no permitido",
			http.StatusMethodNotAllowed)
	}
}

func tratamientosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createTratamiento(w, r)
	default:
		http.Error(w, "Method not allowed",
			http.StatusMethodNotAllowed)
	}
}
