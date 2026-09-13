package main

import (
	sqlc "TPE/db/sqlc"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func getVoucher(w http.ResponseWriter, r *http.Request, id int) {

	voucher, err := queries.GetVoucher(ctx, int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(voucher)
}

func deleteVoucher(w http.ResponseWriter, r *http.Request, id int) {

	err := queries.DeleteVoucher(ctx, int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func createVoucher(w http.ResponseWriter, r *http.Request) {
	var nuevoVoucher sqlc.Voucher
	err := json.NewDecoder(r.Body).Decode(&nuevoVoucher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	paramsCreacion := sqlc.CreateVoucherParams{RegaladorID: nuevoVoucher.RegaladorID, TratamientoID: nuevoVoucher.TratamientoID, ClienteID: nuevoVoucher.ClienteID}
	nuevoVoucherRow, _ := queries.CreateVoucher(ctx, paramsCreacion)
	json.NewEncoder(w).Encode(nuevoVoucherRow)
}

func updateVoucher(w http.ResponseWriter, r *http.Request, id int) {
	var voucherToUpdate sqlc.Voucher
	err := json.NewDecoder(r.Body).Decode(&voucherToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	voucherToUpdateParams := sqlc.UpdateVoucherParams{IDVoucher: int64(id), RegaladorID: voucherToUpdate.RegaladorID, TratamientoID: voucherToUpdate.TratamientoID, ClienteID: voucherToUpdate.ClienteID}
	err = queries.UpdateVoucher(ctx, voucherToUpdateParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	updatedVoucher, _ := queries.GetVoucher(ctx, voucherToUpdate.ClienteID)
	json.NewEncoder(w).Encode(updatedVoucher)
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

func vouchersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createVoucher(w, r)
	default:
		http.Error(w, "Method not allowed",
			http.StatusMethodNotAllowed)
	}
}
