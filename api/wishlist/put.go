package wishlist

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kalmecak/bucherliste/common"
)

// Put Modifica el contenido de una wishlist
func Put(c *fiber.Ctx) error {

	id := c.Context().UserValue("id").(*uuid.UUID)
	res := common.Response{}
	res.Ok("Updated: " + id.String())
	return c.Status(fiber.StatusOK).JSON(&res)
}
