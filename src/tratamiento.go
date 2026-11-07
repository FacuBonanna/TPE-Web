package main

import (
	sqlc "TPE/db/sqlc"
	"context"
	"net/http"
	"strconv"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func getTratamiento(ctx context.Context) []sqlc.Tratamiento {
	tratamientos, _ := queries.ListTreatments(ctx)
	return tratamientos
}

func crearTratamiento(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	nombre := r.FormValue("nombre")
	descripcionCortaStr := r.FormValue("descripcion_corta")
	costoStr := r.FormValue("costo")

	costo, err := strconv.ParseFloat(costoStr, 64)
	if err != nil {
		http.Error(w, "El costo debe ser un número válido", http.StatusBadRequest)
		return
	}

	_, err = queries.CreateTreatment(r.Context(), sqlc.CreateTreatmentParams{
		Nombre:           nombre,
		DescripcionCorta: descripcionCortaStr,
		Costo:            (int32(costo)),
	})
	if err != nil {
		http.Error(w, "Error al crear el tratamiento: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
