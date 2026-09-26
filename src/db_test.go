package main

import (
	sqlc "TPE/db/sqlc"
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	pgtype "github.com/jackc/pgx/v5/pgtype"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestQueries_CRUD(t *testing.T) {

	dbString := "host=localhost port=5432 user=postgres password=postgres dbname=apirest sslmode=disable"
	db, err := sql.Open("pgx", dbString)
	if err != nil {
		t.Fatalf("No se pudo abrir la DB local para el test: %v", err)
	}
	if err := db.Ping(); err != nil {
		t.Fatalf("La DB de Docker no responde en localhost:5432. ¿Está el puerto expuesto? Error: %v", err)
	}
	defer db.Close()

	queries = sqlc.New(db)
	ctx = context.Background()

	var clienteID int64
	var otroClienteID int64
	var tratamientoID int64
	var voucherID int64

	// CLIENTE

	t.Run("Crear cliente", func(t *testing.T) {
		paramsCreacion := sqlc.CreateUserParams{Nombre: "valentina", Apellido: "bisogni", Deuda: 10000, NroTelefono: 123456}

		cliente, err := queries.CreateUser(ctx, paramsCreacion)
		if err != nil {
			t.Errorf("Error crítico al crear el cliente en la DB: %v", err)
		} else {
			fmt.Println("Se creó exitosamente el cliente.")
			clienteID = cliente.ID
		}
	})

	t.Run("Get cliente", func(t *testing.T) {
		cliente, err := queries.GetUser(ctx, clienteID)
		if err != nil {
			t.Errorf("No se puedo obtener el cliente: %v", err)
		} else {
			if cliente.Nombre != "valentina" || cliente.Apellido != "bisogni" || cliente.Deuda != 10000 || cliente.NroTelefono != 123456 {
				t.Errorf("No se pudo obtener el cliente: %v", err)
			} else {
				fmt.Println("El cliente se obtuvo con éxito.")
			}
		}
	})

	t.Run("Update cliente", func(t *testing.T) {

		paramsClienteToUpdate := sqlc.UpdateUserParams{ID: clienteID, Nombre: "valentina", Apellido: "bisogni", Deuda: 5000, NroTelefono: 123456}
		_, err := queries.UpdateUser(ctx, paramsClienteToUpdate)
		if err != nil {
			t.Errorf("No se puedo actualizar el cliente: %v", err)
		} else {
			fmt.Println("El cliente se actualizó exitosamente.")
		}
	})

	t.Run("Comprobar update cliente", func(t *testing.T) {
		cliente, err := queries.GetUser(ctx, clienteID)

		if err != nil || cliente.Nombre != "valentina" || cliente.Apellido != "bisogni" || cliente.Deuda != 5000 || cliente.NroTelefono != 123456 {
			t.Errorf("no se pudo obtener el cliente actualizado: %v", err)
		} else {
			fmt.Println("Se comprobó la actualización de cliente")
		}
	})

	//TRATAMIENTO

	t.Run("Crear tratamiento", func(t *testing.T) {
		paramsCreacion := sqlc.CreateTreatmentParams{Nombre: "masaje deportivo", DescripcionCorta: "masajes", Costo: 10000}
		tratamiento, err := queries.CreateTreatment(ctx, paramsCreacion)
		if err != nil {
			t.Errorf("No se creó el tratamiento:")
		} else {
			fmt.Println("El tratamiento se creó con éxito.")
			tratamientoID = tratamiento.ID
		}
	})

	t.Run("Get Tratamiento", func(t *testing.T) {
		tratamiento, err := queries.GetTreatment(ctx, tratamientoID)
		if err != nil {
			t.Errorf("No se pudo obtener el Tratameinto: %v", err)
		} else {
			if tratamiento.Nombre != "masaje deportivo" || tratamiento.DescripcionCorta != "masajes" || tratamiento.Costo != 10000 {
				t.Errorf("No se pudo obtener el tratamiento: %v", err)
			} else {
				fmt.Println("El tratamiento fue recuperado con éxito.")
			}
		}
	})

	t.Run("Actualizar Tratamiento", func(t *testing.T) {
		updateParams := sqlc.UpdateTreatmentParams{ID: tratamientoID, Nombre: "masajes deportivos", DescripcionCorta: "masajes", Costo: 12000}
		_, err := queries.UpdateTreatment(ctx, updateParams)
		if err != nil {
			t.Errorf("Error crítico al intentar actualizar el tratamiento en la DB: %v", err)
		} else {
			fmt.Println("El tratamiento se actualizó exitosamente.")
		}
	})

	t.Run("Comprobar modificación tratamiento", func(t *testing.T) {
		tratamiento, err := queries.GetTreatment(ctx, tratamientoID)

		if err != nil {
			t.Errorf("Error crítico al intentar actualizar el voucher en la DB: %v", err)
		} else {
			if int64(tratamiento.Costo) == 12000 {
				fmt.Println("Se comprobó la actualización del tratamiento.")
			} else {
				t.Error("No se pudo comporbar la actualización del tratamiento.")
			}
		}
	})

	//VOUCHER

	t.Run("Crear voucher", func(t *testing.T) {
		paramsCreacionUsuario := sqlc.CreateUserParams{Nombre: "Facundo", Apellido: "Bonanna", Deuda: 5000, NroTelefono: 123456}
		user, err := queries.CreateUser(ctx, paramsCreacionUsuario)
		otroClienteID = user.ID
		paramsCreacion := sqlc.CreateVoucherParams{RegaladorID: clienteID, ClienteID: otroClienteID, TratamientoID: 1}

		voucher, err := queries.CreateVoucher(ctx, paramsCreacion)
		if err != nil {
			t.Errorf("Error crítico al crear el voucher en la DB: %v", err)
		} else {
			fmt.Println("El voucher se creó exitosamente.")
			voucherID = voucher.IDVoucher
		}
	})

	t.Run("Get voucher", func(t *testing.T) {

		queries.GetVoucher(ctx, voucherID)
		if err != nil {
			t.Errorf("Error crítico al intentar obtener el voucher en la DB: %v", err)
		} else {
			fmt.Println("Se obtuvo exitosamente el voucher.")

		}
	})

	t.Run("Modificar voucher", func(t *testing.T) {
		updateParams := sqlc.UpdateVoucherParams{IDVoucher: voucherID, ClienteID: clienteID, RegaladorID: clienteID, TratamientoID: tratamientoID}
		_, err := queries.UpdateVoucher(ctx, updateParams)
		if err != nil {
			t.Errorf("Error crítico al intentar actualizar el voucher en la DB: %v", err)
		} else {
			fmt.Println("El voucher se actualizó exitosamente.")
		}
	})

	t.Run("Comprobar modificación voucher", func(t *testing.T) {
		voucher, err := queries.GetVoucher(ctx, voucherID)

		if err != nil {
			t.Errorf("Error crítico al intentar actualizar el voucher en la DB: %v", err)
		} else {
			if voucher.ClienteID != otroClienteID {
				fmt.Println("Se comprobó la actualización del voucher.")
			} else {
				t.Error("No se pudo comprobar la actualización del voucher.")
			}
		}
	})

	t.Run("Eliminar voucher", func(t *testing.T) {
		err := queries.DeleteVoucher(ctx, voucherID)

		if err != nil {
			t.Errorf("No se pudo eliminar el voucher: %v", err)
		} else {
			fmt.Println("Se eliminó el voucher.")
		}
	})

	t.Run("Comprobar eliminación voucher", func(t *testing.T) {
		_, err := queries.GetVoucher(ctx, clienteID)

		if err != nil {
			fmt.Println("Se comprobó la eliminación del voucher.")
		} else {
			t.Errorf("Falló la comprobación de eliminación del voucher :%v", err)
		}

	})

	//TURNO

	var fecha pgtype.Date
	var hora int32

	fecha = pgtype.Date{
		Time:  time.Date(2026, 9, 25, 0, 0, 0, 0, time.UTC),
		Valid: true,
	}
	hora = 1

	t.Run("Crear turno", func(t *testing.T) {
		paramsCreacion := sqlc.CreateTurnoParams{Fecha: fecha, Hora: hora, TratamientoID: tratamientoID, ClienteID: clienteID}
		_, err := queries.CreateTurno(ctx, paramsCreacion)
		if err != nil {
			t.Errorf("No se pudo crear un turno :%V", err)
		} else {
			fmt.Println("El turno se creó con exito.")
		}
	})

	t.Run("Get turno", func(t *testing.T) {
		paramsGet := sqlc.GetTurnoParams{Fecha: fecha, Hora: hora}
		turno, err := queries.GetTurno(ctx, paramsGet)

		if err != nil {
			t.Errorf("No se pudo obtener el turno :%v", err)
		} else {
			if turno.Fecha != fecha || turno.Hora != hora || turno.TratamientoID != tratamientoID || turno.ClienteID != clienteID {
				fmt.Print("aca")
				t.Errorf("Turno no se pudo obtener :%v", err)
			} else {
				fmt.Println("Turno se obtuvo exitosamente.")
			}
		}

	})

	t.Run("Update Turno", func(t *testing.T) {
		turnoToUpdate := sqlc.UpdateTurnoParams{Fecha: fecha, Hora: hora, TratamientoID: tratamientoID, ClienteID: otroClienteID}
		_, err := queries.UpdateTurno(ctx, turnoToUpdate)

		if err != nil {
			t.Errorf("No se pudo actualizar el turno: %v", err)
		} else {
			fmt.Println("El turno se actualizó exitosamente.")
		}
	})

	t.Run("Comprobar actualizacion turno", func(t *testing.T) {
		paramsGet := sqlc.GetTurnoParams{Fecha: fecha, Hora: hora}
		turno, err := queries.GetTurno(ctx, paramsGet)

		if err != nil {
			t.Errorf("Turno no se pudo obtener luego de actualizacion :%v", err)
		} else {
			if turno.Fecha != fecha || turno.Hora != hora || turno.TratamientoID != tratamientoID || turno.ClienteID != otroClienteID {
				t.Errorf("Turno no se pudo obtener luego de actualizar por parametros :%v", err)
			} else {
				fmt.Println("se comprobó la actualización de los turnos.")
			}
		}

	})

	t.Run("Eliminar turno", func(t *testing.T) {
		turnoABorrar := sqlc.DeleteTurnoParams{Fecha: fecha, Hora: hora}
		err := queries.DeleteTurno(ctx, turnoABorrar)
		if err != nil {
			t.Errorf("No se pudo eliminar el turno: %v", err)
		} else {
			fmt.Print("EL turno se eliminó exitosamente.")
		}
	})

	t.Run("Comprobar eliminación del turno", func(t *testing.T) {
		paramsGet := sqlc.GetTurnoParams{Fecha: fecha, Hora: hora}
		_, err := queries.GetTurno(ctx, paramsGet)

		if err != nil {
			fmt.Println("Se comprobó la eliminacion del turno")
		} else {
			t.Errorf("No se comprobó la eliminacion del turno: %v", err)
		}
	})

	//DELETE CLIENTE Y TRATAMIENTO

	t.Run("Eliminar cliente", func(t *testing.T) {
		err := queries.DeleteUser(ctx, clienteID)

		if err != nil {
			t.Errorf("No se pudo eliminar el cliente: %v", err)
		} else {
			fmt.Println("Se eliminó el cliente.")
		}
	})

	t.Run("Comprobar eliminación cliente", func(t *testing.T) {
		_, err := queries.GetUser(ctx, clienteID)

		if err != nil {
			fmt.Println("Se comprobó la eliminación del cliente.")
		} else {
			t.Errorf("No se pudo comprobar la eliminacion del cliente :%v", err)
		}

	})

	t.Run("Eliminar tratamiento", func(t *testing.T) {
		err := queries.DeleteTreatment(ctx, tratamientoID)
		if err != nil {
			t.Errorf("No se pudo eliminar el tratamiento: %v", err)
		} else {
			fmt.Println("Se eliminó el tratamiento.")
		}
	})

	t.Run("Comprobar eliminación tratamiento", func(t *testing.T) {
		_, err := queries.GetTreatment(ctx, tratamientoID)

		if err != nil {
			fmt.Println("Se comprobó la eliminación del tratamiento.")
		} else {
			t.Errorf("No se pudo comprobar la eliminación del tratamiento :%v", err)
		}

	})

}
