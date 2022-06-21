package common

import (
	"os"
	"strconv"
)

// jwtTTL es el jwt time to live, lee le variable de entorno JWT_TTL y devuelve el valor
// en minutos, por default es de 15 nimutos
var jwtTTL = func() int {

	m := os.Getenv("JWT_TTL")
	if m == "" {

		return 15
	}

	if mins, err := strconv.Atoi(m); err != nil {

		return 15
	} else {

		return mins
	}
}()
