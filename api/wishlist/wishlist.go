package wishlist

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
	"github.com/kalmecak/bucherliste/sql"
	logger "github.com/kalmecak/go-error-logger"
	"gorm.io/gorm"
)

// Wishlist responde el contenido de una wishlist
func Wishlist(c *fiber.Ctx) error {

	id := c.Context().UserValue("paramID").(*sql.UID)
	user := c.Context().UserValue("userID").(*sql.UID)
	wl := &sql.Wishlist{
		ID: *id,
	}

	// Conectamos con la BD
	db, err := sql.GormDB()
	if err != nil {

		// Log de error y respueta de error interno
		logger.Error(err, "api.wishlist.Books.db.GormDB")
		var res common.Response
		res.InternalError("could not connect to BD", "")
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}
	// Buscamos la wishlist en la BD
	tx := db.Where("user_id = ?", user).First(wl)
	if err := tx.Error; err != nil {

		res := booksRes{}

		// si el error es NotFound, respondemos con una wishlist vacía
		if err == gorm.ErrRecordNotFound {

			res.NotFound("wishlist not found")
			c.Status(fiber.StatusNotFound)
		} else {

			// Si es otro error, log de error y respuesta de error interno
			logger.Error(err, "api.wishlist.Books.db.First")
			res.InternalError("could not find wishlist", "")
			c.Status(fiber.StatusInternalServerError)
		}
		return c.JSON(&res)
	}

	// Buscamos los libros de la wishlist
	books := []sql.Book{}
	err = db.Model(wl).Association("Books").Find(&books)
	if err != nil {

		logger.Error(err, "api.wishlist.Books.db.Association")
		res := booksRes{}
		res.InternalError("could not find wishlist related books", "")
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(&res)
	}

	res := booksRes{
		Books: books,
	}
	res.Ok("")
	return c.Status(fiber.StatusOK).JSON(&res)
}

type booksRes struct {
	common.Response
	Books []sql.Book `json:"books"`
}
