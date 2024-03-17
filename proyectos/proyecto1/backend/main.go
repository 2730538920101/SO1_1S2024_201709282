// main.go

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"backend/db"
	"backend/manejadores"

	"backend/common"

	"github.com/gorilla/handlers"
)

func main() {
	// Cargar variables de entorno desde el archivo .env
	err := common.CargarVariablesEntorno()
	if err != nil {
		fmt.Println("Error cargando variables de entorno desde el archivo .env:", err)
		fmt.Println("Obteniendo variables de entorno del sistema...")
	}

	// Inicializar la conexión a la base de datos
	dbConnectionString := obtenerDBConnectionString()
	db.InitDB(dbConnectionString)

	defer db.CloseDB() // Asegurar que la conexión a la base de datos se cierre al final

	go manejadores.ActualizarDatosCPU() // Iniciar la rutina de actualización de datos de CPU
	go manejadores.ActualizarDatosRAM() // Iniciar la rutina de actualización de datos de RAM

	router := manejadores.NewRouter()
	port := obtenerPuerto() // Obtener el puerto desde las variables de entorno o establecer uno predeterminado

	fmt.Printf("Escuchando en el puerto %d", port)

	// Usa el número del puerto en la llamada a ListenAndServe
	err = http.ListenAndServe(fmt.Sprintf(":%d", port),
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)(router))
	if err != nil {
		panic(err)
	}
}

// Función para obtener el puerto desde las variables de entorno o establecer uno predeterminado
func obtenerPuerto() int {
	portStr := os.Getenv("SERVER_PORT")
	if portStr == "" {
		// Si la variable de entorno PORT no está configurada, usa el puerto 5000 como predeterminado
		return 5000
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		// Manejar el error en caso de que el valor de la variable de entorno no sea un número
		fmt.Println("Error al convertir el valor de PORT a un número:", err)
		// Usa el puerto 5000 como predeterminado en caso de error
		return 5000
	}

	return port
}

// Función para obtener la cadena de conexión a la base de datos desde las variables de entorno
func obtenerDBConnectionString() string {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	// Componer la cadena de conexión a la base de datos
	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	return dbConnectionString
}
