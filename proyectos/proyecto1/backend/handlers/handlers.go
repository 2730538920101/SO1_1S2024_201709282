package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter devuelve un enrutador configurado con manejadores
func ApiTest() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¡Hola, mundo!")
	}).Methods("GET")

	return router
}
