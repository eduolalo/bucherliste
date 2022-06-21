package validate

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
	"github.com/kalmecak/bucherliste/sql"
	logger "github.com/kalmecak/go-error-logger"
)

// IDParam valida que el id enviado en los id son uuid.UUID
func IDParam(c *fiber.Ctx) error {

	uid, err := sql.UIDFromString(c.Params("id"))
	if err != nil {

		logger.Error(err, "api.validate.IDParam.id.Scan")
		var res common.Response
		res.BadRequest("Invalid ID")
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	if uid.String() == "00000000-0000-0000-0000-000000000000" {

		logger.Message("Id vacío", "api.validate.IDParam.id.String")
		var res common.Response
		res.BadRequest("Empty ID")
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}
	c.Context().SetUserValue("paramID", &uid)
	return c.Next()
}
