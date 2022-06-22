package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
	"github.com/kalmecak/bucherliste/sql"
	logger "github.com/kalmecak/go-error-logger"
	"gorm.io/gorm"
)

// LogIn Genera un token de autenticación en caso de que las credenciales recibidas
// pertenezcan a un usuario registrado
func LogIn(c *fiber.Ctx) error {

	// tomamos la estrcutura de usuario almacenada en el contexto
	user := c.Context().UserValue("user").(*sql.User)
	// Conectamos con la BD
	db, err := sql.GormDB()
	if err != nil {

		// Log de error y respueta de error interno
		logger.Error(err, "api.user.login.db.GormDB")
		var res common.Response
		res.InternalError("could not connect to BD", "")
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}
	if err := user.BuilHash(); err != nil {

		// Log de error y respueta de error interno
		logger.Error(err, "api.user.login.user.BuildHash")
		var res logged
		res.InternalError("could not parse user password", "")
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}
	// Buscamos el usuario en la BD
	tx := db.Where("username = ?", user.Username).
		Where("hash = ?", user.Hash).First(user)

	if err := tx.Error; err != nil {

		// Revisamos si encontró el usuario
		var res logged
		if err == gorm.ErrRecordNotFound {

			res.NotFound("user not found")
			c.Status(fiber.StatusNotFound)
		} else {

			// Si es otro error, log de error y respuesta de error interno
			logger.Error(err, "api.user.logIn.db.First")
			res.InternalError("could not find user", "")
			c.Status(fiber.StatusInternalServerError)
		}
		return c.JSON(&res)
	}

	// Generamos el token
	payload := common.Payload{
		Ref: user.ID.String(),
	}

	token, err := payload.JWT()
	if err != nil {

		// Log de error y respueta de error interno
		logger.Error(err, "api.user.login.payload.JWT")
		var res logged
		res.InternalError("could not create jwt", "")
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}
	res := logged{
		Token: token,
	}
	res.Ok("")
	return c.Status(fiber.StatusOK).JSON(&res)
}

// logged Estructura de la respuesta de login
type logged struct {
	common.Response
	Token string `json:"access_token"`
}
