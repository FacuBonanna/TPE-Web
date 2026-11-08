package main

import (
	"context"
	"net/http"
	"strconv"

	sqlc "TPE/db/sqlc"
	"TPE/views"
)

func getClientes(ctx context.Context) []sqlc.Cliente {
	cli, _ := queries.ListUsers(ctx)
	return cli
}

func getCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id_cliente"), 10, 64)
	if err != nil {
		http.Error(w, "El ID del cliente debe ser un número entero válido", http.StatusBadRequest)
		return
	}

	cliente, err := queries.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, "No se encontró un cliente con ese ID", http.StatusNotFound)
		return
	}

	clientes := []sqlc.Cliente{cliente}
	tratamientos := getTratamiento(r.Context())
	turnos := getTurno(r.Context())

	err = views.Index(clientes, tratamientos, turnos).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error al renderizar la página", http.StatusInternalServerError)
	}
}

func crearCliente(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	nombre := r.FormValue("name")
	apellido := r.FormValue("surname")

	// Convertir deuda a float64
	deuda, err := strconv.ParseFloat(r.FormValue("debt"), 64)
	if err != nil {
		http.Error(w, "La deuda debe ser un número válido", http.StatusBadRequest)
		return
	}

	// Convertir teléfono a int32
	telefono, err := strconv.ParseInt(r.FormValue("phone"), 10, 32)
	if err != nil {
		http.Error(w, "El teléfono debe ser un número entero válido", http.StatusBadRequest)
		return
	}

	_, err = queries.CreateUser(r.Context(), sqlc.CreateUserParams{
		Nombre:      nombre,
		Apellido:    apellido,
		Deuda:       deuda,
		NroTelefono: int32(telefono),
	})

	if err != nil {
		http.Error(w, "Error al crear el cliente: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
