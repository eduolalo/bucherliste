package validate

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
	"github.com/kalmecak/bucherliste/sql"
	logger "github.com/kalmecak/go-error-logger"
)

// CreateBody valida que el body para crear una wishlist sea correcto
func CreateBody(c *fiber.Ctx) error {

	var wl sql.Wishlist
	// parseamos el body del request dentro del wl struct
	if err := wl.Unmarshal(c.Body()); err != nil {

		logger.Error(err, "api.validate.CreateBody.wl.Unmarshal")
		var res common.Response
		res.InternalError("error parsing body", "")
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}

	// Validamos que las propiedades del wl cumplan con los requisitos de la BD
	if err := wl.Validate(); err != nil {

		logger.Error(err, "api.validate.CreateBody.wl.Validate")
		var res common.Response
		res.BadRequest(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(&res)
	}

	c.Context().SetUserValue("wl", &wl)
	return c.Next()
}
