package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Security maneja los middleware de seguridad y para optimizar las peticiones
func Security(app *fiber.App) {

	// Configuración de CORS
	app.Use(cors.New(cors.Config{
		// Permite todos los orígenes
		AllowOrigins: "*",
		// *****************************************************************************
		// De cuerdo con las DISPOSICIONES de carácter general relativas a las API's   *
		// informáticas estandarizadas a que hace referencia la Ley para Regular las   *
		// ITF's, ANEXO 1, numeral 2.1, punto II,                                      *
		// insizo "a":                                                                 *
		// SÓLO SE PODRÁ TRABAJAR CON LOS VERVOS GET, POST, PUT y DELETE               *
		// *****************************************************************************
		// Sólo se deben agregar los vervos que se usan en el API
		AllowMethods: "GET,POST,PUT,DELETE",
		// Optional. Default value []string{}.
		AllowHeaders: "", //"Content-Type, Authorization",
		// Optional. Default value false.
		// AllowCredentials bool,
		// Optional. Default value []string{}.
		// ExposeHeaders []string,
		// Optional. Default value 0.
		// MaxAge int,
	}))

}
