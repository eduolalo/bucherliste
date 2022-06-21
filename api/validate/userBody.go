package validate

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
	"github.com/kalmecak/bucherliste/sql"
	logger "github.com/kalmecak/go-error-logger"
)

// UserBody revisa que el body del request contenga los datos necesarios para el registro
// y loggeo de un usuario
func UserBody(c *fiber.Ctx) error {

	var user sql.User
	// parseamos el body del request dentro del user struct
	if err := user.Unmarshal(c.Body()); err != nil {

		logger.Error(err, "api.validate.signUpBody.user.Unmarshal")
		var res common.Response
		res.InternalError("error parsing body", "")
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}

	// Validamos que las propiedades del user cumplan con los requisitos de la BD
	if err := user.Validate(); err != nil {

		logger.Error(err, "api.validate.signUpBody.user.Validate")
		var res common.Response
		res.BadRequest(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(&res)
	}

	c.Context().SetUserValue("user", &user)
	return c.Next()
}
