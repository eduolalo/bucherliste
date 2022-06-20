package wishlist

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
)

// Delete Elimina una wishlist
func Delete(c *fiber.Ctx) error {

	id := c.Params("id")
	res := common.Response{}
	res.Ok("deleted: " + id)
	return c.Status(fiber.StatusOK).JSON(&res)
}
