package wishlist

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
	"github.com/kalmecak/bucherliste/sql"
)

// Get Obtiene todas las wishlists del usuario
func Get(c *fiber.Ctx) error {

	res := listRes{
		Wishlists: []*sql.Wishlist{},
	}
	res.Ok("")
	return c.Status(fiber.StatusOK).JSON(&res)
}

type listRes struct {
	common.Response
	Wishlists []*sql.Wishlist `json:"wishlists"`
}
