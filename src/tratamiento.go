package main

import (
	sqlc "TPE/db/sqlc"
	"TPE/views"
	"context"
	"net/http"
	"strconv"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func getTratamientos(ctx context.Context) []sqlc.Tratamiento {
	tratamientos, _ := queries.ListTreatments(ctx)
	return tratamientos
}

func listarTratamientos(w http.ResponseWriter, r *http.Request) {
	views.FilasTratamientos(getTratamientos(r.Context())).Render(r.Context(), w)
}

func getTratamiento(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id_tratamiento"), 10, 64)
	if err != nil {
		views.FilasTratamientos(nil).Render(r.Context(), w)
		return
	}

	tratamiento, err := queries.GetTreatment(r.Context(), id)
	if err != nil {
		views.FilasTratamientos(nil).Render(r.Context(), w)
		return
	}

	tratamientos := []sqlc.Tratamiento{tratamiento}

	err = views.FilasTratamientos(tratamientos).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error al renderizar la página", http.StatusInternalServerError)
	}
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

	tratamientos := getTratamientos(r.Context())
	err = views.TablaTratamientos(tratamientos).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error al renderizar los tratamientos", http.StatusInternalServerError)
		return
	}
}

func deleteTratamiento(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "El ID del tratamiento debe ser un número entero válido", http.StatusBadRequest)
		return
	}

	err = queries.DeleteTreatment(r.Context(), id)
	if err != nil {
		http.Error(w, "Error al eliminar el tratamiento: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getTratamientoFila(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	t, err := queries.GetTreatment(r.Context(), id)
	if err != nil {
		http.Error(w, "Tratamiento no encontrado", http.StatusNotFound)
		return
	}

	_ = views.FilaTratamiento(t).Render(r.Context(), w)
}

func editarTratamientoFila(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	t, err := queries.GetTreatment(r.Context(), id)
	if err != nil {
		http.Error(w, "Tratamiento no encontrado", http.StatusNotFound)
		return
	}

	_ = views.FilaTratamientoEdit(t).Render(r.Context(), w)
}

func actualizarTratamiento(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	nombre := r.FormValue("nombre")
	descripcionCorta := r.FormValue("descripcion_corta")
	costo, err := strconv.ParseFloat(r.FormValue("costo"), 64)
	if err != nil {
		http.Error(w, "Costo inválido", http.StatusBadRequest)
		return
	}

	t, err := queries.UpdateTreatment(r.Context(), sqlc.UpdateTreatmentParams{
		ID:               id,
		Nombre:           nombre,
		DescripcionCorta: descripcionCorta,
		Costo:            int32(costo),
	})
	if err != nil {
		http.Error(w, "Error al actualizar el tratamiento: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_ = views.FilaTratamiento(t).Render(r.Context(), w)
}
