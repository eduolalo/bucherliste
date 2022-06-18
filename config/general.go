package config

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// General maneja los middleware de seguridad y para optimizar las peticiones
func General(app *fiber.App) {

	// Usar la máxima compresión
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// Log de los requests
	format := "--> ${method} ${path} <--\n ${time} - HTTPCode: ${status} - IP: ${ips} ${ip} - Latencia: ${latency}\n"
	app.Use(logger.New(logger.Config{
		Format:     format,
		TimeFormat: "02-Jan-2006",
	}))

	if os.Getenv("DEBUG") == "true" {

		app.Use("/", debugLog)
	}

	// Revisión de las cabeceras necesarias
	// app.Use("/user", checkHeaders)
	// app.Use("/deposit", checkHeaders)
	// app.Use("/dispersion", checkHeaders)
}

// checkHeaders revisa las cabeceras de los paths especidicados
// func checkHeaders(c *fiber.Ctx) error {

// 	// Content-type sea de tipo JSON
// 	c.Type("json", "utf-8")

// 	// Que la cabecera de autorización esté presente en todos los requests
// 	fsctx := c.Context()
// 	header := string(fsctx.Request.Header.Peek("Authorization"))
// 	if header == "" {

// 		elogger.Message("Cabecera 'Authorization' faltante..", "config.General.checkHeaders")
// 		var res common.Response
// 		res.Forbridden("Cabecera 'Authorization' faltante.")
// 		return c.Status(fiber.StatusForbidden).JSON(res)
// 	}

// 	// Cabecera W-Body-Sign de firma para los PUT/POST
// 	if c.Method() == "POST" || c.Method() == "PUT" {

// 		if header = string(fsctx.Request.Header.Peek("W-Body-Sign")); header == "" {

// 			elogger.Message("Cabecera 'W-Body-Sign' faltante.", "config.General.checkHeaders")
// 			var res common.Response
// 			res.Forbridden("Cabecera 'W-Body-Sign' faltante.")
// 			return c.Status(fiber.StatusForbidden).JSON(res)
// 		}
// 	}

// 	return c.Next()

// }

// debugLog tira los datos de la solicitud
func debugLog(ctx *fiber.Ctx) error {

	log.Println("*** debug request ***")

	if ctx.Request() != nil {

		headers := ctx.Request().Header.Header()
		log.Println("headers: \n", string(headers))
	}

	log.Println("Ruta: ", ctx.OriginalURL())
	log.Println("IP: ", ctx.IP())
	log.Println("Body: ", string(ctx.Body()))
	log.Println("Content-Type JSON: ", ctx.Is("json"))

	log.Println("--- debug request ---")

	return ctx.Next()
}
