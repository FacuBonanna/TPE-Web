package main

import (
sqlc "TPE/db/sqlc"
	"context"
	"database/sql"
	"fmt"
	"testing"

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

	t.Run ("Eliminar cliente", func( t *testing.T){
		err := queries.DeleteUser(ctx, clienteID)

		if err != nil {
			t.Errorf("No se pudo eliminar el cliente: %v", err)
		} else {
			fmt.Print("Se elimino el cliente")
		}
	})

	t.Run ("comprobar eliminacion", func(t *testing.T){
		//get del eliminado para comprobar que se elimino
		_, err := queries.GetUser(ctx, clienteID)

		if err != nil{
			fmt.Print("Se comprobo la eliminacion del cliente")
		} else {
			t.Errorf("fallo la comprobacion de eliminacion del cliente :%v", err)
		}

	})
}

