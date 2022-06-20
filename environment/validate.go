package environment

import (
	"log"
	"os"
)

// Validate - Analiza que las variables de entorno esten configuradas
// correctamente y se encuentren en el .profile del contenedor
func Validate() (ok bool) {

	variables := []string{
		"DB_STRING",
		"JWT_SECRET",
		"GO_ENV",
		"LOG_LEVEL",
		"DEBUG",
	}
	counter := 0
	for k := range variables {

		val := os.Getenv(variables[k])
		if val != "" {
			counter++
		} else {

			log.Println(variables[k] + ": no encontrado.")
		}
	}

	return (counter == len(variables))
}
