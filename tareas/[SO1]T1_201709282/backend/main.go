package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Message struct {
    Nombre string `json:"nombre"`
    Carnet  int    `json:"carnet"`
}

// Controlador que maneja las solicitudes a la ruta "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
    // Crear un objeto Message
    message := Message{Nombre: "CARLOS JAVIER MARTINEZ POLANCO", Carnet: 201709282,}
    // Convertir el objeto a JSON
    jsonResponse, err := json.Marshal(message)
    if err != nil {
        http.Error(w, "Error al serializar el objeto JSON", http.StatusInternalServerError)
        return
    }

    // Establecer el encabezado de respuesta
    w.Header().Set("Content-Type", "application/json")

    // Escribir la respuesta
    w.Write(jsonResponse)
}

func main() {
    router := mux.NewRouter()
    // Configurar el manejador para la ruta "/"
	router.HandleFunc("/", homeHandler)

    corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://t1_frontend:80", "http://localhost"}, // Reemplaza con el origen de tu aplicación React
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(router)

    // Iniciar el servidor en el puerto 5000
    err := http.ListenAndServe(":5000", corsHandler)
    if err != nil {
        panic(err)
    }
}
