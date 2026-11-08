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

func getTratamiento(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("tratamiento_id"), 10, 64)
	if err != nil {
		http.Error(w, "El ID del tratamiento debe ser un número entero válido", http.StatusBadRequest)
		return
	}

	tratamiento, err := queries.GetTreatment(r.Context(), id)
	if err != nil {
		http.Error(w, "No se encontró un tratamiento con ese Documento", http.StatusNotFound)
		return
	}

	clientes := getClientes(r.Context())
	tratamientos := []sqlc.Tratamiento{tratamiento}
	turnos := getTurnos(r.Context())
	voucher := getVouchers(r.Context())

	err = views.Index(clientes, tratamientos, turnos, voucher).Render(r.Context(), w)
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

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
