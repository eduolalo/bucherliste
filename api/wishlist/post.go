package wishlist

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
)

// Post Crea una nueva wishlist ligada al usuario
func Post(c *fiber.Ctx) error {

	res := common.Response{}
	res.Ok("")
	return c.Status(fiber.StatusOK).JSON(&res)
}
