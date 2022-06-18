package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
	logger "github.com/kalmecak/go-error-logger"
)

// Page500 es el manejador de respuestas para cuando haya erores internos
func Page500(c *fiber.Ctx, err error) error {

	logger.Error(err, "config.Page500")
	var response common.Response
	response.InternalError("uups, sorry ¯\\_(ツ)_/¯", "")
	return c.Status(fiber.StatusInternalServerError).JSON(response)
}
