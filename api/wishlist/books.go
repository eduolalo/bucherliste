package wishlist

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
	"github.com/kalmecak/bucherliste/sql"
)

// Books responde el contenido de una wishlist
func Books(c *fiber.Ctx) error {

	res := booksRes{
		Wishlist: &sql.Wishlist{
			Books: []sql.Book{},
		},
	}
	res.Ok("")
	return c.Status(fiber.StatusOK).JSON(&res)
}

type booksRes struct {
	common.Response
	Wishlist *sql.Wishlist `json:"wishlist"`
}
