// db/db.go

package db

import (
	"database/sql"
	"fmt"

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
