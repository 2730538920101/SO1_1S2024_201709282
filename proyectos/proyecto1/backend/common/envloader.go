package common

import (
	"fmt"

	"github.com/joho/godotenv"
)

// CargarVariablesEntorno carga las variables de entorno desde el archivo .env
func CargarVariablesEntorno() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error al cargar variables de entorno: %w", err)
	}
	return nil
}
