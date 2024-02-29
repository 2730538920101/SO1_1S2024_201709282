package handlers

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

const (
	cpuFilePath = "/proc/cpu_so1_1s2024" // Ruta al archivo de datos de CPU
	ramFilePath = "/proc/ram_so1_1s2024" // Ruta al archivo de datos de RAM
)

var cpuDataChan = make(chan string) // Canal para datos de CPU
var ramDataChan = make(chan string) // Canal para datos de RAM

// ActualizarDatosCPU obtiene datos de CPU desde el archivo en /proc y los envía al canal
func ActualizarDatosCPU() {
	for {
		datosCPU, err := ObtenerDatosDesdeArchivo(cpuFilePath)
		if err != nil {
			fmt.Println("Error al obtener datos de CPU:", err)
			// Manejar el error según sea necesario
		}
		cpuDataChan <- datosCPU

		// Actualizar cada minuto, puedes ajustar el intervalo según tus necesidades
		time.Sleep(time.Minute)
	}
}

// ActualizarDatosRAM obtiene datos de RAM desde el archivo en /proc y los envía al canal
func ActualizarDatosRAM() {
	for {
		datosRAM, err := ObtenerDatosDesdeArchivo(ramFilePath)
		if err != nil {
			fmt.Println("Error al obtener datos de RAM:", err)
			// Manejar el error según sea necesario
		}
		ramDataChan <- datosRAM

		// Actualizar cada minuto, puedes ajustar el intervalo según tus necesidades
		time.Sleep(time.Minute)
	}
}

// ObtenerDatosDesdeArchivo ejecuta un cat al archivo y devuelve su contenido
func ObtenerDatosDesdeArchivo(filePath string) (string, error) {
	cmd := exec.Command("cat", filePath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error: %s", err.Error())
	}

	return strings.TrimSpace(string(out)), nil
}

// HandleCPUDatos retorna los datos de CPU al endpoint correspondiente
func HandleCPUDatos(w http.ResponseWriter, r *http.Request) {
	datosCPU := <-cpuDataChan
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"informacion_cpu": %s}`, datosCPU)
}

// HandleRAMDatos retorna los datos de RAM al endpoint correspondiente
func HandleRAMDatos(w http.ResponseWriter, r *http.Request) {
	datosRAM := <-ramDataChan
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `%s`, datosRAM)
}

// NewRouter devuelve un enrutador configurado con manejadores
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Endpoint para datos de CPU
	router.HandleFunc("/cpu", HandleCPUDatos).Methods("GET")

	// Endpoint para datos de RAM
	router.HandleFunc("/ram", HandleRAMDatos).Methods("GET")

	return router
}
