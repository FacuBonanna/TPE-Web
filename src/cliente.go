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

func getCliente(w http.ResponseWriter, r *http.Request, id int) {

	cliente, err := queries.GetUser(ctx, int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(cliente)
}

func deleteCliente(w http.ResponseWriter, r *http.Request, id int) {

	err := queries.DeleteUser(ctx, int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func updateCliente(w http.ResponseWriter, r *http.Request, id int) {
	var clienteToUpdate sqlc.Cliente
	err := json.NewDecoder(r.Body).Decode(&clienteToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	paramsClienteToUpdate := sqlc.UpdateUserParams{ID: int64(id), Nombre: clienteToUpdate.Nombre, Apellido: clienteToUpdate.Apellido, Deuda: clienteToUpdate.Deuda, NroTelefono: clienteToUpdate.NroTelefono}
	updatedCliente, err := queries.UpdateUser(ctx, paramsClienteToUpdate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedCliente)
}

func createCliente(w http.ResponseWriter, r *http.Request) {
	var nuevoCliente sqlc.Cliente
	err := json.NewDecoder(r.Body).Decode(&nuevoCliente)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Print(err.Error())
		return
	}
	paramsCreacion := sqlc.CreateUserParams{Nombre: nuevoCliente.Nombre, Apellido: nuevoCliente.Apellido, Deuda: nuevoCliente.Deuda, NroTelefono: nuevoCliente.NroTelefono}
	nuevoClienteRow, err := queries.CreateUser(ctx, paramsCreacion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(nuevoClienteRow)
	}
}

func clientesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createCliente(w, r)
	default:
		http.Error(w, "Method not allowed",
			http.StatusMethodNotAllowed)
	}
}

func clienteHandler(w http.ResponseWriter, r *http.Request) {
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
		getCliente(w, r, id)
	case http.MethodPut:
		updateCliente(w, r, id)
	case http.MethodDelete:
		deleteCliente(w, r, id)
	default:
		http.Error(w, "Método no permitido",
			http.StatusMethodNotAllowed)
	}
}
