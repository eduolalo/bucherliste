package books

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
)

// Get obtiene una lista de libros según los parámetros de búsqueda
func Get(c *fiber.Ctx) error {
	res := common.Response{}
	res.Ok("")
	return c.Status(fiber.StatusOK).JSON(&res)
}
