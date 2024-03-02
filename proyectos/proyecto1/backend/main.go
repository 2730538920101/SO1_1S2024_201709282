// main.go

package main

import (
	"fmt"
	"net/http"

	"backend/db"
	"backend/handlers"
)

func main() {
	// Inicializar la conexión a la base de datos
	db.InitDB("sopes:sopes1234@tcp(localhost:3306)/proyecto1")

	defer db.CloseDB() // Asegurar que la conexión a la base de datos se cierre al final

	go handlers.ActualizarDatosCPU() // Iniciar la rutina de actualización de datos de CPU
	go handlers.ActualizarDatosRAM() // Iniciar la rutina de actualización de datos de RAM

	router := handlers.NewRouter()
	port := 5000
	fmt.Printf("Escuchando en http://localhost:%d\n", port)

	// Usa el número del puerto en la llamada a ListenAndServe
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		panic(err)
	}
}
