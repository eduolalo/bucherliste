package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
	logger "github.com/kalmecak/go-error-logger"
)

// Page404 maneja las peticiones cuyo path no son encontrados
func Page404(app *fiber.App) {

	app.Use(func(c *fiber.Ctx) error {

		mssg := "Página no encontrada." + c.Path()
		logger.Message(mssg, "config.Page404")
		var res common.Response
		res.NotFound("Check your request path >:v")
		return c.Status(fiber.StatusNotFound).JSON(res)
	})
}
