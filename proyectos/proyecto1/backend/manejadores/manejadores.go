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

// GenerarArbolDOT genera el árbol en formato DOT para el PID dado
func GenerarArbolDOT(pid string) (string, error) {
	// Obtén los datos de CPU desde el canal
	datosCPU := <-cpuDataChan

	// Deserializa los datos JSON en una estructura models.InformacionProcesos
	var informacionProcesos models.InformacionProcesos
	if err := json.Unmarshal([]byte(datosCPU), &informacionProcesos); err != nil {
		return "", fmt.Errorf("error al deserializar datos de CPU: %w", err)
	}

	// Busca el proceso con el PID específico en la lista de procesos
	var procesoSeleccionado *models.ProcesoPadre
	for _, proceso := range informacionProcesos.Procesos {
		if fmt.Sprintf("%d", proceso.PID) == pid {
			procesoSeleccionado = &proceso
			break
		}
	}

	// Si no se encontró el proceso, devuelve un mensaje indicando que no se encontró
	if procesoSeleccionado == nil {
		return "", fmt.Errorf("proceso con PID %s no encontrado", pid)
	}

	// Genera el árbol DOT utilizando la información del proceso seleccionado
	arbolDot := generateDOTTree(procesoSeleccionado)

	// Agrega el encabezado 'digraph' al árbol DOT
	arbolDot = "digraph G {\n" + arbolDot + "}\n"

	return arbolDot, nil
}

// Función auxiliar para generar el árbol DOT recursivamente
func generateDOTTree(proceso *models.ProcesoPadre) string {
	// Estructura básica del nodo para el proceso actual
	nodeString := fmt.Sprintf("%d [label=\"%s\"];\n", proceso.PID, proceso.Nombre)

	// Agrega conexiones con los procesos hijos
	for _, hijo := range proceso.Hijos {
		nodeString += generateDOTTreeHijo(hijo)
		nodeString += fmt.Sprintf("%d -> %d;\n", proceso.PID, hijo.PIDHijo)
	}

	return nodeString
}

// Función auxiliar para generar el árbol DOT para un proceso hijo recursivamente
func generateDOTTreeHijo(hijo models.ProcesoHijo) string {
	// Estructura básica del nodo para el proceso hijo actual
	nodeString := fmt.Sprintf("%d [label=\"%s\"];\n", hijo.PIDHijo, hijo.NombreHijo)

	// No se agregan conexiones a los procesos hijos, ya que son hojas en el árbol

	return nodeString
}

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

// HandleCPUDatosLista retorna la lista de datos de CPU al endpoint correspondiente
func HandleCPUDatosLista(w http.ResponseWriter, r *http.Request) {
	listaCPU, err := db.ObtenerListaCPUUltimos10Minutos()
	if err != nil {
		http.Error(w, fmt.Sprintf("error al obtener lista de datos de CPU: %s", err), http.StatusInternalServerError)
		return
	}

	// Crear la estructura de respuesta
	respuestaCPU := models.CPUDatosResponse{
		ListaCPU: listaCPU,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuestaCPU)
}

// HandleRAMDatosLista retorna la lista de datos de RAM al endpoint correspondiente
func HandleRAMDatosLista(w http.ResponseWriter, r *http.Request) {
	listaRAM, err := db.ObtenerListaRAMUltimos10Minutos()
	if err != nil {
		http.Error(w, fmt.Sprintf("error al obtener lista de datos de RAM: %s", err), http.StatusInternalServerError)
		return
	}

	// Crear la estructura de respuesta
	respuestaRAM := models.RAMDatosResponse{
		ListaRAM: listaRAM,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuestaRAM)
}

// HandleListaPIDProcesos retorna la lista de PID de procesos al endpoint correspondiente
func HandleListaPIDProcesos(w http.ResponseWriter, r *http.Request) {
	// Obtener los datos de CPU desde el canal
	datosCPU := <-cpuDataChan

	// Deserializar los datos JSON en una estructura models.InformacionProcesos
	var informacionProcesos models.InformacionProcesos
	if err := json.Unmarshal([]byte(datosCPU), &informacionProcesos); err != nil {
		http.Error(w, fmt.Sprintf("error al deserializar datos de CPU: %s", err), http.StatusInternalServerError)
		return
	}

	// Utilizar un mapa para verificar duplicados
	pidMap := make(map[int]bool)
	var listaPID []int

	// Obtener la lista de PID de procesos
	for _, proceso := range informacionProcesos.Procesos {
		// Verificar si el PID ya existe en el mapa antes de agregarlo
		if _, exists := pidMap[proceso.PID]; !exists {
			pidMap[proceso.PID] = true
			listaPID = append(listaPID, proceso.PID)
		}
	}

	// Crear un mapa para la respuesta
	respuesta := map[string]interface{}{
		"pids": listaPID,
	}

	// Convertir la respuesta a formato JSON
	jsonData, err := json.Marshal(respuesta)
	if err != nil {
		http.Error(w, fmt.Sprintf("error al serializar lista de PID: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func HandleGenerarArbol(w http.ResponseWriter, r *http.Request) {
	// Obtén el PID del parámetro de la URL
	vars := mux.Vars(r)
	pid, ok := vars["pid"]
	if !ok {
		http.Error(w, "Falta el parámetro 'pid' en la URL", http.StatusBadRequest)
		return
	}

	// Llama a una función para generar el árbol en formato DOT usando el PID
	arbolDot, err := GenerarArbolDOT(pid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al generar el árbol DOT: %s", err), http.StatusInternalServerError)
		return
	}

	// Devuelve el árbol en formato DOT como respuesta
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, arbolDot)
}

// NewRouter devuelve un enrutador configurado con manejadores
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Endpoint para datos de CPU
	router.HandleFunc("/cpu", HandleCPUDatos).Methods("GET")

	// Endpoint para datos de RAM
	router.HandleFunc("/ram", HandleRAMDatos).Methods("GET")

	// Endpoint para historico de CPU
	router.HandleFunc("/historico_cpu", HandleCPUDatosLista).Methods("GET")

	// Endpoint para historico de la RAM
	router.HandleFunc("/historico_ram", HandleRAMDatosLista).Methods("GET")

	// Endpoint para obtener la lista de PID de procesos
	router.HandleFunc("/lista_pid_procesos", HandleListaPIDProcesos).Methods("GET")

	// Endpoint para obtener el arbol del proceso seleccionado
	router.HandleFunc("/generarArbol/{pid}", HandleGenerarArbol).Methods("GET")

	return router
}
