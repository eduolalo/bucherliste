package config

import (
	"strings"

	"github.com/kalmecak/bucherliste/common"
	logger "github.com/kalmecak/go-error-logger"

	"github.com/gofiber/fiber/v2"
)

// AuthHeader ejecuta las reviciones de la cabacera de autorización, en caso que venga
// un JWT, se revisa la validéz y se estrae la información
func AuthHeader(app *fiber.App) {

	app.Use(func(c *fiber.Ctx) error {

		fsctx := c.Context()
		header := string(fsctx.Request.Header.Peek("Authorization"))

		// Si la cabecera de autorización está vacía, se continua con el flujo ya que no
		// todos los paths necesitan que el usuario tenga sesión
		if header == "" {

			c.Context().SetUserValue("payload", common.Payload{})
			c.Context().SetUserValue("isLogged", false)
			return c.Next()
		}

		// Se busca el estandar oauth2 en la cabecera
		if !strings.HasPrefix(header, "Bearer ") {

			logger.Message("Authorization header is not Bearer", "config.AuthHeader.HasPrefix")
			var res common.Response
			res.BadRequest("Authorization header is not Bearer")
			return c.Status(fiber.StatusBadRequest).JSON(res)
		}

		// Se divide el string de la cabecera en un array de strings buscar el jwt
		sliced := strings.Split(header, " ")
		// Revisamos que sliced sea de longitud 2
		if len(sliced) != 2 {

			logger.Message("Authorization header malformed", "config.AuthHeader.Header.len != 2")
			var res common.Response
			res.BadRequest("Authorization header malformed")
			return c.Status(fiber.StatusBadRequest).JSON(res)
		}

		// Validación del JWT e integración de la información en la estructura
		payload := common.Payload{}
		if err := payload.Validate(sliced[1]); err != nil {

			logger.Message("jwt corrumted", "config.AuthHeader.payload.Validate")
			var res common.Response
			res.Forbridden("Token corrupted")
			return c.Status(fiber.StatusForbidden).JSON(res)
		}

		// se revisa que el payload tenga los datos necesarios
		if err := payload.IsValid(); err != nil {

			logger.Message("payload malformed", "config.AuthHeader.payload.IsValid")
			var res common.Response
			res.Forbridden("payload malformed")
			return c.Status(fiber.StatusForbidden).JSON(res)
		}
		c.Context().SetUserValue("payload", &payload)
		c.Context().SetUserValue("isLogged", true)
		return c.Next()
	})
}
