// models/models.go

package models

import "sync"

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

// Mutex para proteger la sección crítica
var Mutex sync.Mutex
