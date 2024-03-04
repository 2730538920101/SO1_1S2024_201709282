// models/models.go

package models

import (
	"sync"
	"time"
)

// Estructura para representar un proceso hijo
type ProcesoHijo struct {
	IDProcesoHijo int    `json:"ID_PROCESO_HIJO"`
	PIDHijo       int    `json:"PID_HIJO"`
	NombreHijo    string `json:"Nombre_HIJO"`
	EstadoHijo    int    `json:"Estado_HIJO"`
	RSSHijo       int    `json:"RSS_HIJO"`
	UIDHijo       int    `json:"UID_HIJO"`
}

// Estructura para representar un proceso padre
type ProcesoPadre struct {
	IDProceso int           `json:"ID_PROCESO"`
	PID       int           `json:"PID"`
	Nombre    string        `json:"Nombre"`
	Estado    int           `json:"Estado"`
	RSS       int           `json:"RSS"`
	UID       int           `json:"UID"`
	Hijos     []ProcesoHijo `json:"hijos"`
}

// Estructura principal
type InformacionProcesos struct {
	PorcentajeCPU float64        `json:"porcentaje_utilizacion_cpu"`
	Procesos      []ProcesoPadre `json:"procesos"`
}

// RAM representa la estructura de información de RAM
type RAM struct {
	InformacionMemoria struct {
		TotalMemoria        int     `json:"total_memoria"`
		MemoriaLibre        int     `json:"memoria_libre"`
		MemoriaUtilizada    int     `json:"memoria_utilizada"`
		PorcentajeUtilizado float64 `json:"porcentaje_utilizado"`
	} `json:"informacion_memoria"`
}

// CPUData representa la estructura para almacenar datos de CPU en la base de datos
type CPUData struct {
	ID             int       `json:"id"`
	Porcentaje     float64   `json:"porcentaje"`
	TiempoRegistro time.Time `json:"tiempo_registro"`
}

// RAMData representa la estructura para almacenar datos de RAM en la base de datos
type RAMData struct {
	ID                  int       `json:"id"`
	PorcentajeUtilizado float64   `json:"porcentaje_utilizado"`
	TiempoRegistro      time.Time `json:"tiempo_registro"`
}

// CPUDatosResponse representa el formato de respuesta deseado para datos de CPU
type CPUDatosResponse struct {
	ListaCPU []CPUData `json:"lista_cpu"`
}

// RAMDatosResponse representa el formato de respuesta deseado para datos de RAM
type RAMDatosResponse struct {
	ListaRAM []RAMData `json:"lista_ram"`
}

// Mutex para proteger la sección crítica
var Mutex sync.Mutex
