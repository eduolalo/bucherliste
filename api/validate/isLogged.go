package validate

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
	logger "github.com/kalmecak/go-error-logger"
)

// IsLogged Revisa que el usuario esté logueado
func IsLogged(c *fiber.Ctx) error {

	logged := c.Context().UserValue("isLogged").(*bool)
	log.Println("IsLogged: ", *logged)
	// si está logueado, se continua con el flujo
	if *logged {

		return c.Next()
	}

	// si no está logueado, se devuelve un Forbidden
	logger.Message("Request sin sesión", "api.validate.IsLogged")
	var res common.Response
	res.Forbridden("session expired")
	return c.Status(fiber.StatusForbidden).JSON(res)
}
