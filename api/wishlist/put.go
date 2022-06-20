package wishlist

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
)

// Put Modifica el contenido de una wishlist
func Put(c *fiber.Ctx) error {

	id := c.Params("id")
	res := common.Response{}
	res.Ok("Updated: " + id)
	return c.Status(fiber.StatusOK).JSON(&res)
}
