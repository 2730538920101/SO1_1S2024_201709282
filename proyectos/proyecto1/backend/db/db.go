// db/db.go

package db

import (
	"backend/models"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // Importa el controlador MySl
)

var db *sql.DB

// InitDB inicializa la conexión a la base de datos
func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	// Comprueba si la conexión a la base de datos es exitosa
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Conexión a la base de datos exitosa")
}

// CloseDB cierra la conexión a la base de datos
func CloseDB() {
	db.Close()
}

// InsertProceso inserta datos de un proceso en la base de datos y devuelve el ID generado
func InsertProceso(pid, nombre string, estado, rss, uid int) (int, error) {
	result, err := db.Exec("INSERT INTO PROCESO (PID, NOMBRE, ESTADO, RSS, UID) VALUES (?, ?, ?, ?, ?)", pid, nombre, estado, rss, uid)
	if err != nil {
		return 0, err
	}

	idProceso, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(idProceso), nil
}

// InsertProcesoHijo inserta datos de un proceso hijo en la base de datos
func InsertProcesoHijo(idProceso int, pid, nombre string, estado, rss, uid int) error {
	_, err := db.Exec("INSERT INTO PROCESO_HIJO (ID_PROCESO, PID, NOMBRE, ESTADO, RSS, UID) VALUES (?, ?, ?, ?, ?, ?)", idProceso, pid, nombre, estado, rss, uid)
	return err
}

// InsertRAM inserta datos de RAM en la base de datos
func InsertRAM(totalMemoria, memoriaLibre, memoriaUtilizada int, porcentajeUtilizado float64) error {
	_, err := db.Exec("INSERT INTO RAM (TOTAL_MEMORIA, MEMORIA_LIBRE, MEMORIA_UTILIZADA, PORCENTAJE_UTILIZADO) VALUES (?, ?, ?, ?)",
		totalMemoria, memoriaLibre, memoriaUtilizada, porcentajeUtilizado)
	return err
}

// InsertCPU inserta datos de CPU en la base de datos
func InsertCPU(porcentajeUtilizado float64) error {
	_, err := db.Exec("INSERT INTO CPU (PORCENTAJE_UTILIZADO) VALUES (?)",
		porcentajeUtilizado)
	return err
}

// ObtenerListaCPUUltimos10Minutos obtiene la lista de porcentajes de CPU y sus marcas de tiempo
func ObtenerListaCPUUltimos10Minutos() ([]models.CPUData, error) {
	tiempoLimite := time.Now().Add(-10 * time.Minute)
	rows, err := db.Query("SELECT ID_CPU, PORCENTAJE_UTILIZADO, TIEMPO FROM CPU WHERE TIEMPO >= ? ORDER BY TIEMPO", tiempoLimite)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listaCPU []models.CPUData
	for rows.Next() {
		var data models.CPUData
		var tiempoDB []uint8
		if err := rows.Scan(&data.ID, &data.Porcentaje, &tiempoDB); err != nil {
			return nil, err
		}

		// Convierte []uint8 a time.Time
		tiempo, err := time.Parse("2006-01-02 15:04:05", string(tiempoDB))
		if err != nil {
			return nil, err
		}

		data.TiempoRegistro = tiempo
		listaCPU = append(listaCPU, data)
	}

	return listaCPU, nil
}

// ObtenerListaRAMUltimos10Minutos obtiene la lista de porcentajes de RAM y sus marcas de tiempo
func ObtenerListaRAMUltimos10Minutos() ([]models.RAMData, error) {
	tiempoLimite := time.Now().Add(-10 * time.Minute)
	rows, err := db.Query("SELECT ID_RAM, PORCENTAJE_UTILIZADO, TIEMPO FROM RAM WHERE TIEMPO >= ? ORDER BY TIEMPO", tiempoLimite)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listaRAM []models.RAMData
	for rows.Next() {
		var data models.RAMData
		var tiempoDB []uint8
		if err := rows.Scan(&data.ID, &data.PorcentajeUtilizado, &tiempoDB); err != nil {
			return nil, err
		}

		// Convierte []uint8 a time.Time
		tiempo, err := time.Parse("2006-01-02 15:04:05", string(tiempoDB))
		if err != nil {
			return nil, err
		}

		data.TiempoRegistro = tiempo
		listaRAM = append(listaRAM, data)
	}

	return listaRAM, nil
}
