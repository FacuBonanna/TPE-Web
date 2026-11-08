package main

import (
	sqlc "TPE/db/sqlc"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func getVoucher(w http.ResponseWriter, r *http.Request, id int) {

	voucher, err := queries.GetVoucher(r.Context(), int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(voucher)
}

func deleteVoucher(w http.ResponseWriter, r *http.Request, id int) {

	err := queries.DeleteVoucher(r.Context(), int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func createVoucher(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	// Convertir documentos a int32
	regalador_documento, err := strconv.ParseInt(r.FormValue("regalador_documento"), 10, 32)
	if err != nil {
		http.Error(w, "El documento debe ser un número entero válido", http.StatusBadRequest)
		return
	}
	cliente_documento, err := strconv.ParseInt(r.FormValue("cliente_documento"), 10, 32)
	if err != nil {
		http.Error(w, "El documento debe ser un número entero válido", http.StatusBadRequest)
		return
	}
	tratamiento_id, err := strconv.ParseInt(r.FormValue("tratamiento_id"), 10, 32)
	if err != nil {
		http.Error(w, "El ID del tratamiento debe ser un número entero válido", http.StatusBadRequest)
		return
	}

	_, err = queries.CreateVoucher(r.Context(), sqlc.CreateVoucherParams{
		RegaladorID:   regalador_documento,
		ClienteID:     cliente_documento,
		TratamientoID: tratamiento_id,
	})

	if err != nil {
		http.Error(w, "Error al crear el voucher: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func updateVoucher(w http.ResponseWriter, r *http.Request, id int) {
	var voucherToUpdate sqlc.Voucher
	err := json.NewDecoder(r.Body).Decode(&voucherToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	voucherToUpdateParams := sqlc.UpdateVoucherParams{IDVoucher: int64(id), RegaladorID: voucherToUpdate.RegaladorID, TratamientoID: voucherToUpdate.TratamientoID, ClienteID: voucherToUpdate.ClienteID}
	updatedVoucher, err := queries.UpdateVoucher(r.Context(), voucherToUpdateParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedVoucher)
}

func getVouchers(ctx context.Context) []sqlc.Voucher {
	vouchers, _ := queries.ListVouchers(ctx)
	return vouchers
}

func voucherHandler(w http.ResponseWriter, r *http.Request) {
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
		getVoucher(w, r, id)
	case http.MethodPut:
		updateVoucher(w, r, id)
	case http.MethodDelete:
		deleteVoucher(w, r, id)
	default:
		http.Error(w, "Método no permitido",
			http.StatusMethodNotAllowed)
	}
}
