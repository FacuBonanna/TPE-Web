package main

import (
	sqlc "TPE/db/sqlc"
	"TPE/views"
	"context"
	"net/http"
	"strconv"
)

func getVoucher(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("voucher_id"), 10, 64)
	if err != nil {
		http.Error(w, "El ID del voucher debe ser un número entero válido", http.StatusBadRequest)
		return
	}

	voucher, err := queries.GetVoucher(r.Context(), id)
	if err != nil {
		http.Error(w, "No se encontró un voucher con ese ID", http.StatusNotFound)
		return
	}

	clientes := getClientes(r.Context())
	tratamientos := getTratamientos(r.Context())
	turnos := getTurnos(r.Context())
	vouchers := sqlc.GetVouchersRow{IDVoucher: voucher.IDVoucher, RegaladorDocumento: voucher.RegaladorDocumento, ClienteDocumento: voucher.ClienteDocumento, TratamientoID: voucher.TratamientoID, Regalador: voucher.Regalador, Cliente: voucher.Cliente, Tratamiento: voucher.Tratamiento}
	vouchersArr := []sqlc.GetVouchersRow{vouchers}

	err = views.Index(clientes, tratamientos, turnos, vouchersArr).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error al renderizar la página", http.StatusInternalServerError)
	}
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
		RegaladorDocumento: regalador_documento,
		ClienteDocumento:   cliente_documento,
		TratamientoID:      tratamiento_id,
	})

	if err != nil {
		http.Error(w, "Error al crear el voucher: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func getVouchers(ctx context.Context) []sqlc.GetVouchersRow {
	vouchers, _ := queries.GetVouchers(ctx)
	return vouchers
}
