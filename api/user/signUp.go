package user

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
	"github.com/kalmecak/bucherliste/sql"
	logger "github.com/kalmecak/go-error-logger"
)

func SignUp(c *fiber.Ctx) error {

	// tomamos la estrcutura de usuario almacenada en el contexto
	user := c.Context().UserValue("user").(*sql.User)

	// Conectamos con la BD
	db, err := sql.GormDB()
	if err != nil {

		// Log de error y respueta de error interno
		logger.Error(err, "api.user.signUp.db.GormDB")
		var res common.Response
		res.InternalError("could not connect to BD", "")
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}

	// Creamos el usuario en la BD
	tx := db.Create(user)
	if err := tx.Error; err != nil {

		var res common.Response
		// Manejamos si ya está registrado un el username
		if strings.Contains(err.Error(), "Duplicate entry") {

			res.BadRequest("user already exists")
			c.Status(fiber.StatusBadRequest)
		} else {

			// Si no está registrado, log de error y respuesta de error interno
			logger.Error(err, "api.user.signUp.db.Create")
			res.InternalError("could not create user", "")
			c.Status(fiber.StatusInternalServerError)
		}
		return c.JSON(&res)
	}

	res := common.Response{}
	res.Created("")
	return c.Status(fiber.StatusCreated).JSON(&res)
}
