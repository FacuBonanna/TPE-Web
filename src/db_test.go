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


//cliente 
func TestQueriesCliente_CRUD(t *testing.T) {
	// Debido a que go corre en la propia maquina 
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

	t.Run ("Crear cliente", func(t *testing.T) {  
	paramsCreacion := sqlc.CreateUserParams{Nombre: "valentina", Apellido: "bisogni", Deuda: 10000, NroTelefono: 123456}
	
	cliente, err := queries.CreateUser(ctx, paramsCreacion)
	if err != nil { 
		t.Errorf("Error crítico al crear el cliente en la DB: %v", err)
	} else { 
		fmt.Println("se creo exitosamente el cliente")
		clienteID = cliente.ID
	}})

	t.Run("get cliente", func(t *testing.T) {
		cliente, err := queries.GetUser(ctx, clienteID)
		if err != nil {
			t.Errorf("No se puedo obtener el cliente: %v", err)
		} else {
			if cliente.Nombre != "valentina" || cliente.Apellido != "bisogni" || cliente.Deuda != 10000 || cliente.NroTelefono != 123456 {
				t.Errorf("no se pudo obtener el cliente: %v", err)
			} else { 
				fmt.Print("Se obtuvo con exito el cliente")}
		}
	})

	t.Run ("Update cliente", func(t *testing.T){
		//cambio la dueda
		paramsClienteToUpdate := sqlc.UpdateUserParams{ID: clienteID, Nombre: "valentina", Apellido: "bisogni", Deuda: 5000, NroTelefono: 123456}
		_, err := queries.UpdateUser(ctx, paramsClienteToUpdate)
		if err != nil {
			t.Errorf("No se puedo actualizar el cliente: %v", err)
		} else {
				fmt.Print("Se actualizo exitosamente al cliente")
			}
	})

	t.Run ("Comprobar update", func(t *testing.T){
		cliente, err := queries.GetUser(ctx, clienteID)
			//esta bien mezclar logica del error con esto??
			if err != nil || cliente.Nombre != "valentina" || cliente.Apellido != "bisogni" || cliente.Deuda != 5000 || cliente.NroTelefono != 123456 {
				t.Errorf("no se pudo obtener el cliente actualizado: %v", err)
			} else {
				fmt.Print("Se comprobo la actualizacion de cliente")
			}
	})


	//Tratamiento 
	var tratamientoID int64

	t.Run ("Crear tratamiento", func(t *testing.T) {  
		paramsCreacion := sqlc.CreateTreatmentParams{Nombre: "masaje deportivo", DescripcionCorta: "masajes", Costo: 10000}
		tratamiento, err := queries.CreateTreatment(ctx, paramsCreacion)	
		if err != nil {
			t.Errorf("No se creó el tratamiento:")
		} else { fmt.Print("El tratamiento se creó con éxito")
			tratamientoID = tratamiento.ID
			_ = tratamientoID
	}})



	//TURNO
	var fecha pgtype.Date
	var hora int32 

	fecha = pgtype.Date{
		Time : time.Date(2026,9,25,0,0,0,0, time.UTC),
		Valid : true,
	}
	hora = 1

	t.Run ("Crear turno", func(t *testing.T){
		paramsCreacion := sqlc.CreateTurnoParams{Fecha: fecha, Hora: hora, TratamientoID: tratamientoID, ClienteID: clienteID}
		_, err := queries.CreateTurno(ctx, paramsCreacion)
		if err != nil {
			t.Errorf("No se pudo crear un turno :%V", err)
		} else {
			fmt.Print("El turno se creo con exito")
		}
	})

		t.Run ("Get turno", func(t *testing.T){
		paramsGet := sqlc.GetTurnoParams{Fecha: fecha, Hora: hora}
		turno, err := queries.GetTurno(ctx, paramsGet)

		if err != nil {
			t.Errorf("Turno no se pudo obtener :%v", err)
		} else {
			if turno.Fecha != fecha || turno.Hora != hora || turno.TratamientoID != tratamientoID || turno.ClienteID != clienteID{
				fmt.Print("aca")
				t.Errorf("Turno no se pudo obtener :%v", err)
			} else {
				fmt.Print("Turno se obtuvo exitosamente")
			}
		}

	})

	t.Run ("Update Turno", func(t *testing.T){
		//cambio hora a 2
		turnoToUpdate := sqlc.UpdateTurnoParams{Fecha: fecha, Hora: hora, TratamientoID: tratamientoID, ClienteID: clienteID}
		_, err := queries.UpdateTurno(ctx,turnoToUpdate)

		if err != nil {
			t.Errorf("No se pudo actualizar el turno: %v", err)
		} else {
			fmt.Print("el turno se actualizo exitosamente")
		}

	})

	t.Run ("Eliminar turno", func(t *testing.T){
		turnoABorrar := sqlc.DeleteTurnoParams{Fecha: fecha, Hora: hora}
		err := queries.DeleteTurno(ctx, turnoABorrar)
		if err != nil {
			t.Errorf("No se pudo eliminar el turno: %v", err)
		} else {
			fmt.Print("Se elimino exitosamente el turno")
		}



	})



		t.Run ("Eliminar cliente", func( t *testing.T){
		err := queries.DeleteUser(ctx, clienteID)

		if err != nil {
			t.Errorf("No se pudo eliminar el cliente: %v", err)
		} else {
			fmt.Print("Se elimino el cliente")
		}
	})

	t.Run ("comprobar eliminacion", func(t *testing.T){
		_, err := queries.GetUser(ctx, clienteID)

		if err != nil{
			fmt.Print("Se comprobo la eliminacion del cliente")
		} else {
			t.Errorf("fallo la comprobacion de eliminacion del cliente :%v", err)
		}

	})
	
}