package main

import (
	sqlc "TPE/db/sqlc"
	"TPE/views"
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgconn"
)

func getVoucher(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("voucher_id"), 10, 64)
	if err != nil {
		views.FilasVouchers(nil).Render(r.Context(), w)
		return
	}

	voucher, err := queries.GetVoucher(r.Context(), id)
	if err != nil {
		views.FilasVouchers(nil).Render(r.Context(), w)
		return
	}

	vouchers := sqlc.GetVouchersRow{IDVoucher: voucher.IDVoucher, RegaladorDocumento: voucher.RegaladorDocumento, ClienteDocumento: voucher.ClienteDocumento, TratamientoID: voucher.TratamientoID, Regalador: voucher.Regalador, Cliente: voucher.Cliente, Tratamiento: voucher.Tratamiento}
	vouchersArr := []sqlc.GetVouchersRow{vouchers}

	err = views.FilasVouchers(vouchersArr).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error al renderizar la página", http.StatusInternalServerError)
	}
}

func listarVouchers(w http.ResponseWriter, r *http.Request) {
	views.FilasVouchers(getVouchers(r.Context())).Render(r.Context(), w)
}

func createVoucher(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	// Convertir documentos a int32
	regalador_documento, err := strconv.ParseInt(r.FormValue("regalador_documento"), 10, 32)
	if err != nil {
		errorHTMX(w, r, "El documento del regalador debe ser un número entero válido")
		return
	}
	cliente_documento, err := strconv.ParseInt(r.FormValue("cliente_documento"), 10, 32)
	if err != nil {
		errorHTMX(w, r, "El documento del regalado debe ser un número entero válido")
		return
	}
	tratamiento_id, err := strconv.ParseInt(r.FormValue("tratamiento_id"), 10, 32)
	if err != nil {
		errorHTMX(w, r, "El ID del tratamiento debe ser un número entero válido")
		return
	}

	_, err = queries.CreateVoucher(r.Context(), sqlc.CreateVoucherParams{
		RegaladorDocumento: regalador_documento,
		ClienteDocumento:   cliente_documento,
		TratamientoID:      tratamiento_id,
	})

	if err != nil {
		log.Printf("crearVoucher: %v", err)
		errorHTMX(w, r, getMotivoVoucher(err, "crear"))
		return
	}

	vouchers := getVouchers(r.Context())
	err = views.TablaVouchers(vouchers).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error al renderizar los vouchers", http.StatusInternalServerError)
		return
	}
}

func deleteVoucher(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "El ID del voucher debe ser un número entero válido", http.StatusBadRequest)
		return
	}

	err = queries.DeleteVoucher(r.Context(), id)
	if err != nil {
		http.Error(w, "Error al eliminar el voucher: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getVouchers(ctx context.Context) []sqlc.GetVouchersRow {
	vouchers, _ := queries.GetVouchers(ctx)
	return vouchers
}

func getVoucherFila(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	documento, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	v, err := queries.GetVoucher(r.Context(), documento)
	if err != nil {
		http.Error(w, "Voucher no encontrado", http.StatusNotFound)
		return
	}
	vo := sqlc.GetVouchersRow{IDVoucher: v.IDVoucher, RegaladorDocumento: v.RegaladorDocumento, ClienteDocumento: v.ClienteDocumento, Regalador: v.Regalador, Cliente: v.Cliente, Tratamiento: v.Tratamiento}
	_ = views.FilaVoucher(vo).Render(r.Context(), w)
}

func editarVoucherFila(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	v, err := queries.GetVoucher(r.Context(), id)
	if err != nil {
		http.Error(w, "Voucher no encontrado", http.StatusNotFound)
		return
	}
	sqlcV := sqlc.Voucher{IDVoucher: v.IDVoucher, RegaladorDocumento: v.RegaladorDocumento, TratamientoID: v.TratamientoID, ClienteDocumento: v.ClienteDocumento}
	_ = views.FilaVoucherEdit(sqlcV).Render(r.Context(), w)
}

func actualizarVoucher(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	id := r.PathValue("id_voucher")
	idStr, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "ID de voucher inválido", http.StatusBadRequest)
		return
	}

	documentoRegaladorStr := r.FormValue("regalador_documento")
	documentoRegalador, err := strconv.ParseInt(documentoRegaladorStr, 10, 64)
	if err != nil {
		http.Error(w, "Documento del regalador inválido", http.StatusBadRequest)
		return
	}

	tratamientoId := r.FormValue("tratamiento_id")
	tratamientoIdStr, err := strconv.ParseInt(tratamientoId, 10, 64)
	if err != nil {
		http.Error(w, "ID del tratamiento inválido", http.StatusBadRequest)
		return
	}

	clienteId := r.FormValue("cliente_documento")
	clienteIdStr, err := strconv.ParseInt(clienteId, 10, 64)
	if err != nil {
		http.Error(w, "Documento del cliente inválido", http.StatusBadRequest)
		return
	}

	v, err := queries.UpdateVoucher(r.Context(), sqlc.UpdateVoucherParams{
		IDVoucher:          idStr,
		RegaladorDocumento: documentoRegalador,
		TratamientoID:      tratamientoIdStr,
		ClienteDocumento:   clienteIdStr,
	})
	if err != nil {
		errorHTMX(w, r, getMotivoVoucher(err, "actualizar"))
		return
	}
	vo := sqlc.GetVouchersRow{IDVoucher: v.IDVoucher, RegaladorDocumento: v.RegaladorDocumento, ClienteDocumento: v.ClienteDocumento, Regalador: v.Regalador, Cliente: v.Cliente, Tratamiento: v.Tratamiento}
	_ = views.FilaVoucher(vo).Render(r.Context(), w)
}

func getMotivoVoucher(err error, funcion string) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23503":
			{
				switch pgErr.ConstraintName {
				case "voucher_doc_regalador":
					return "El regalador ingresado no existe."
				case "voucher_cliente":
					return "El cliente ingresado no existe."
				case "voucher_tratamiento":
					return "El tratamiento ingresado no existe."
				}
			}
			return "El cliente o el tratamiento ingresado no existe." + pgErr.ConstraintName
		case "23505":
			return "Ya existe un voucher con esos datos."
		}
	}
	return "No se pudo " + funcion + " el voucher."
}
