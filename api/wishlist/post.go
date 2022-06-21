package wishlist

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
	"github.com/kalmecak/bucherliste/sql"
	logger "github.com/kalmecak/go-error-logger"
)

// Post Crea una nueva wishlist ligada al usuario
func Post(c *fiber.Ctx) error {

	// extraemos la wl y el usuario de la sesion
	id := c.Context().UserValue("userID").(*string)
	wl := c.Context().UserValue("wl").(*sql.Wishlist)

	// Parseamos el id del usuario al tipo UID
	uid, err := sql.UIDFromString(*id)
	if err != nil {

		logger.Error(err, "api.wishlist.Post.wl.UserID.Scan")
		var res common.Response
		res.InternalError("error parsing userID", "")
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}
	// construimos la wishlist
	wl.UserID = uid
	// Conectamos con la BD
	db, err := sql.GormDB()
	if err != nil {

		logger.Error(err, "api.wishlist.Post.db.GormDB")
		var res common.Response
		res.InternalError("could not connect to BD", "")
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}

	// Guardamos la wishlist en la BD
	tx := db.Create(wl)
	if err := tx.Error; err != nil {

		logger.Error(err, "api.wishlist.Post.db.Create")
		var res common.Response
		res.InternalError("could not create wishlist", "")
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}

	if tx.RowsAffected == 0 {

		logger.Message("No se generó la wishlist", "api.wishlist.Post.db.Create")
		var res common.Response
		res.InternalError("could not create wishlist", "")
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}
	res := common.Response{}
	res.Created("")
	return c.Status(fiber.StatusOK).JSON(&res)
}
