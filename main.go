package main

import (
	"fmt"
	"net/http"
)






func main() {
	fileServer := http.FileServer(http.Dir("./"))
	http.Handle("/", fileServer)


	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}
