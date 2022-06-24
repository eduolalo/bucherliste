package main

import (
	"os"

	"github.com/kalmecak/bucherliste/api"
	"github.com/kalmecak/bucherliste/cmd/migration"
	"github.com/kalmecak/bucherliste/config"
	"github.com/kalmecak/bucherliste/environment"

	"log"

	"github.com/gofiber/fiber/v2"
)

var defaultPort = ":8080"

func main() {

	// load azure default port
	if val, ok := os.LookupEnv("PORT"); ok {

		defaultPort = ":" + val
	}

	/**************************************************************************/
	/*               Verifcación de las variables de entorno                  */
	/**************************************************************************/
	if ok := environment.Validate(); !ok {

		log.Panic("Faltan variables de entorno")
	}

	/**************************************************************************/
	/*                        Ejecución del migrate                           */
	/**************************************************************************/

	if err := migration.Start(); err != nil {

		log.Panic(err)
	}

	/**************************************************************************/
	/*                  Creamos la instancia del servidor                     */
	/**************************************************************************/
	app := fiber.New(fiber.Config{
		Prefork:          false,
		ErrorHandler:     config.Page500,
		CaseSensitive:    true,
		StrictRouting:    true,
		AppName:          "Bücherliste",
		DisableKeepalive: true,
	})

	// Routing de los statics

	// Configuraciones de seguridad para el API
	config.Security(app)
	config.General(app)
	config.Accepts(app)
	config.AuthHeader(app)

	// // Routing del API
	api.Router(app)

	// Configuraciones de 404
	config.Page404(app)

	err := app.Listen(defaultPort)
	if err != nil {

		log.Fatal(err)
	}

}
