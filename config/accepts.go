package config

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
	logger "github.com/kalmecak/go-error-logger"
)

// Accepts revisa que el content-type del request sea el correcto
func Accepts(app *fiber.App) {

	app.Use(func(c *fiber.Ctx) error {

		// El content-type sólo se revisa en los métodos que tienen un body
		if c.Method() != "POST" || c.Method() != "PUT" {

			return c.Next()
		}
		// Sólamente se aceptarán los Content-Type appl ication/json
		ct := c.Request().Header.ContentType()
		if !strings.Contains(string(ct), "application/json") {

			logger.Message("Content-Type erroneo.", "config.Accepts")
			var res common.Response
			res.BadRequest("wrong Content-Type")
			return c.Status(fiber.StatusBadRequest).JSON(res)
		}

		return c.Next()
	})
}
