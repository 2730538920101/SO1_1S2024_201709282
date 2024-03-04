package manejadores

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"backend/db"
	"backend/models"

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
			log.Println("Error al obtener datos de CPU:", err)
			// Manejar el error según sea necesario
		}
		cpuDataChan <- datosCPU

		// Actualizar cada 500 milisegundos (0.5 segundos)
		time.Sleep(500 * time.Millisecond)
	}
}

// ActualizarDatosRAM obtiene datos de RAM desde el archivo en /proc y los envía al canal
func ActualizarDatosRAM() {
	for {
		datosRAM, err := ObtenerDatosDesdeArchivo(ramFilePath)
		if err != nil {
			log.Println("Error al obtener datos de RAM:", err)
			// Manejar el error según sea necesario
		}
		ramDataChan <- datosRAM

		// Actualizar cada 500 milisegundos (0.5 segundos)
		time.Sleep(500 * time.Millisecond)
	}
}

// ObtenerDatosDesdeArchivo lee el contenido del archivo y lo devuelve
func ObtenerDatosDesdeArchivo(filePath string) (string, error) {
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error al obtener datos desde el archivo %s: %w", filePath, err)
	}

	return strings.TrimSpace(string(dat)), nil
}

// HandleCPUDatos retorna los datos de CPU al endpoint correspondiente
func HandleCPUDatos(w http.ResponseWriter, r *http.Request) {
	datosCPU := <-cpuDataChan

	// Deserializa los datos JSON en una estructura models.InformacionProcesos
	var informacionProcesos models.InformacionProcesos
	if err := json.Unmarshal([]byte(datosCPU), &informacionProcesos); err != nil {
		http.Error(w, fmt.Sprintf("error al deserializar datos de CPU: %s", err), http.StatusInternalServerError)
		return
	}

	// Mutex para proteger la sección crítica
	models.Mutex.Lock()
	defer models.Mutex.Unlock()
	err := db.InsertCPU(informacionProcesos.PorcentajeCPU)
	if err != nil {
		log.Printf("error al insertar datos de CPU en la base de datos: %s", err)
		http.Error(w, "error interno del servidor", http.StatusInternalServerError)
		return
	}
	// Recorre la lista de procesos y realiza la inserción en la base de datos
	for _, proceso := range informacionProcesos.Procesos {
		// Insertar proceso en la base de datos
		idProceso, err := db.InsertProceso(fmt.Sprintf("%d", proceso.PID), proceso.Nombre, proceso.Estado, proceso.RSS, proceso.UID)
		if err != nil {
			log.Printf("error al insertar datos de CPU en la base de datos: %s", err)
			http.Error(w, "error interno del servidor", http.StatusInternalServerError)
			return
		}

		// Actualiza el ID del proceso en la estructura
		proceso.IDProceso = idProceso

		// Recorre la lista de procesos hijos y realiza la inserción en la base de datos
		for _, hijo := range proceso.Hijos {
			// Insertar proceso hijo en la base de datos
			err := db.InsertProcesoHijo(proceso.IDProceso, fmt.Sprintf("%d", hijo.PIDHijo), hijo.NombreHijo, hijo.EstadoHijo, hijo.RSSHijo, hijo.UIDHijo)
			if err != nil {
				log.Printf("error al insertar datos de CPU (hijo) en la base de datos: %s", err)
				http.Error(w, "error interno del servidor", http.StatusInternalServerError)
				return
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"informacion_cpu": %s}`, datosCPU)
}

// HandleRAMDatos retorna los datos de RAM al endpoint correspondiente
func HandleRAMDatos(w http.ResponseWriter, r *http.Request) {
	datosRAM := <-ramDataChan

	// Deserializa los datos JSON en una estructura models.RAM
	var ram models.RAM
	if err := json.Unmarshal([]byte(datosRAM), &ram); err != nil {
		log.Println("Error al deserializar datos de RAM:", err)
		http.Error(w, "error interno del servidor", http.StatusInternalServerError)
		return
	}

	// Inserta datos en la base de datos
	err := db.InsertRAM(ram.InformacionMemoria.TotalMemoria, ram.InformacionMemoria.MemoriaLibre,
		ram.InformacionMemoria.MemoriaUtilizada, ram.InformacionMemoria.PorcentajeUtilizado)
	if err != nil {
		log.Printf("error al insertar datos de RAM en la base de datos: %s", err)
		http.Error(w, "error interno del servidor", http.StatusInternalServerError)
		return
	}

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
