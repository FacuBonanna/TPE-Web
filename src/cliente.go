package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	sqlc "TPE/db/sqlc"
	"TPE/views"
)

func getClientes(ctx context.Context) []sqlc.Cliente {
	cli, _ := queries.ListUsers(ctx)
	return cli
}

func listarClientes(w http.ResponseWriter, r *http.Request) {
	views.FilasClientes(getClientes(r.Context())).Render(r.Context(), w)
}

func getCliente(w http.ResponseWriter, r *http.Request) {
	documento, err := strconv.ParseInt(r.URL.Query().Get("cliente_documento"), 10, 64)
	if err != nil {
		// 200 a propósito, ver nota abajo
		views.FilasClientes(nil).Render(r.Context(), w)
		return
	}

	cliente, err := queries.GetUser(r.Context(), documento)
	if err != nil {
		views.FilasClientes(nil).Render(r.Context(), w)
		return
	}

	err = views.FilasClientes([]sqlc.Cliente{cliente}).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error al renderizar", http.StatusInternalServerError)
	}
}

func crearCliente(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		errorHTMX(w, r, "Error al procesar el formulario")
		return
	}

	nombre := r.FormValue("name")
	apellido := r.FormValue("surname")

	 if nombre == "" || apellido == "" {
        errorHTMX(w, r, "Faltan campos por completar.")
        return
    }

	// Convertir deuda a float64
	deuda, err := strconv.ParseFloat(r.FormValue("debt"), 64)
	if err != nil {
		errorHTMX(w, r, "La deuda debe ser un número entero válido")
		return
	}

	// Convertir teléfono a int32
	telefono, err := strconv.ParseInt(r.FormValue("phone"), 10, 32)
	if err != nil {
		errorHTMX(w, r, "El teléfono debe ser un número entero válido")
		return
	}

	documento, err := strconv.ParseInt(r.FormValue("cliente_documento"), 10, 32)
	if err != nil {
		errorHTMX(w, r, "El documento debe ser un número entero válido.")
		return
	}

	_, err = queries.CreateUser(r.Context(), sqlc.CreateUserParams{
		Documento:   documento,
		Nombre:      nombre,
		Apellido:    apellido,
		Deuda:       deuda,
		NroTelefono: int32(telefono),
	})

	if err != nil {
		log.Printf("createCliente: %v", err)
		errorHTMX(w, r, motivoCreacion(err, "cliente"))
		return
	}

	clientes := getClientes(r.Context())
	err = views.TablaClientes(clientes).Render(r.Context(), w)
	if err != nil {
		errorHTMX(w, r, "Error al renderizar la lista de clientes")
		return
	}
}

func deleteCliente(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	documento, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		errorHTMX(w, r, "El Documento del cliente debe ser un número entero válido")
		return
	}

	err = queries.DeleteUser(r.Context(), documento)
	if err != nil {
		log.Printf("deleteCliente %d: %v", documento, err) // el detalle técnico queda en el log
		errorHTMX(w, r, motivoBorrado(err, "Cliente"))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getClienteFila(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	documento, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		errorHTMX(w, r, "Documento invalido")
		return
	}

	cliente, err := queries.GetUser(r.Context(), documento)
	if err != nil {
		errorHTMX(w, r, "Cliente no encontrado")
		return
	}

	_ = views.FilaCliente(cliente).Render(r.Context(), w)
}

func editarClienteFila(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	documento, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		errorHTMX(w, r, "Documento invalido")
		return
	}

	cliente, err := queries.GetUser(r.Context(), documento)
	if err != nil {
		errorHTMX(w, r, "Cliente no encontrado")
		return
	}

	_ = views.FilaClienteEdit(cliente).Render(r.Context(), w)
}

func actualizarCliente(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		errorHTMX(w, r, "Error al procesar el formulario")
		return
	}

	idStr := r.PathValue("id")
	documento, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		errorHTMX(w, r, "Documento inválido")
		return
	}

	nombre := r.FormValue("name")
	if nombre == "" {
		errorHTMX(w, r, "El nombre del cliente no puede ser vacío.")
		return
	}
	apellido := r.FormValue("surname")
	if apellido == "" {
		errorHTMX(w, r, "El apellido del cliente no puede ser vacío.")
		return
	}
	deuda, err := strconv.ParseFloat(r.FormValue("debt"), 64)
	if err != nil {
		errorHTMX(w, r, "La deuda debe ser un número válido")
		return
	}
	telefono, err := strconv.ParseInt(r.FormValue("phone"), 10, 32)
	if err != nil {
		errorHTMX(w, r, "El teléfono debe ser un número entero válido")
		return
	}

	cliente, err := queries.UpdateUser(r.Context(), sqlc.UpdateUserParams{
		Documento:   documento,
		Nombre:      nombre,
		Apellido:    apellido,
		Deuda:       deuda,
		NroTelefono: int32(telefono),
	})
	if err != nil {
		errorHTMX(w, r, "Error al actualizar el cliente")
		return
	}

	_ = views.FilaCliente(cliente).Render(r.Context(), w)
}
