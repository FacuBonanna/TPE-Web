package main

import (
	sqlc "TPE/db/sqlc"
	"TPE/views"
	"context"
	"log"
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
		errorHTMX(w, r, "El tratamiento no existe.")
		return
	}
}

func crearTratamiento(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		errorHTMX(w, r, "Error al procesar el formulario")
		return
	}

	nombre := r.FormValue("nombre")
	descripcionCortaStr := r.FormValue("descripcion_corta")
	costoStr := r.FormValue("costo")

	if nombre == "" || descripcionCortaStr == "" {
        errorHTMX(w, r, "Faltan campos por completar.")
        return
    }

	costo, err := strconv.ParseFloat(costoStr, 64)
	if err != nil {
		errorHTMX(w, r, "El costo debe ser un número válido")
		return
	}

	_, err = queries.CreateTreatment(r.Context(), sqlc.CreateTreatmentParams{
		Nombre:           nombre,
		DescripcionCorta: descripcionCortaStr,
		Costo:            (int32(costo)),
	})
	if err != nil {
		log.Printf("createTratamiento: %v", err)
		errorHTMX(w, r, motivoCreacion(err, "tratamiento"))
		return
	}

	tratamientos := getTratamientos(r.Context())
	err = views.TablaTratamientos(tratamientos).Render(r.Context(), w)
	if err != nil {
		errorHTMX(w, r, "Error al renderizar los tratamientos")
		return
	}
}

func deleteTratamiento(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		errorHTMX(w, r, "El ID del tratamiento debe ser un número entero válido")
		return
	}

	err = queries.DeleteTreatment(r.Context(), id)
	if err != nil {
		log.Printf("deleteTreatment %d: %v", id, err) // el detalle técnico queda en el log
		errorHTMX(w, r, motivoBorrado(err,"Tratamiento"))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getTratamientoFila(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		errorHTMX(w, r, "ID invalido")
		return
	}

	t, err := queries.GetTreatment(r.Context(), id)
	if err != nil {
		errorHTMX(w, r, "Tratamiento no encontrado")
		return
	}

	_ = views.FilaTratamiento(t).Render(r.Context(), w)
}

func editarTratamientoFila(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		errorHTMX(w, r, "ID invalido")
		return
	}

	t, err := queries.GetTreatment(r.Context(), id)
	if err != nil {
		errorHTMX(w, r, "Tratamiento no encontrado")
		return
	}

	_ = views.FilaTratamientoEdit(t).Render(r.Context(), w)
}

func actualizarTratamiento(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		errorHTMX(w, r, "Error al procesar el formulario")
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		errorHTMX(w, r, "ID invalido")
		return
	}

	nombre := r.FormValue("nombre")
	if nombre == "" {
		errorHTMX(w, r, "El nombre del tratamiento no puede ser vacío.")
		return
	}

	descripcionCorta := r.FormValue("descripcion_corta")
	if descripcionCorta == "" {
		errorHTMX(w, r, "La del tratamiento no puede ser vacía.")
		return
	}

	costo, err := strconv.ParseFloat(r.FormValue("costo"), 64)
	if err != nil {
		errorHTMX(w, r, "El costo ingresado no es válido.")
		return
	}

	t, err := queries.UpdateTreatment(r.Context(), sqlc.UpdateTreatmentParams{
		ID:               id,
		Nombre:           nombre,
		DescripcionCorta: descripcionCorta,
		Costo:            int32(costo),
	})
	if err != nil {
		errorHTMX(w, r, "Error al actualizar el tratamiento")
		return
	}

	_ = views.FilaTratamiento(t).Render(r.Context(), w)
}
