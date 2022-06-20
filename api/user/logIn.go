package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
)

// LogIn Genera un token de autenticación en caso de que las credenciales recibidas
// pertenezcan a un usuario registrado
func LogIn(c *fiber.Ctx) error {

	res := logged{}
	res.Ok("")
	return c.Status(fiber.StatusOK).JSON(&res)
}

// logged Estructura de la respuesta de login
type logged struct {
	common.Response
	Token string `json:"token"`
}
