package wishlist

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
	"github.com/kalmecak/bucherliste/sql"
	logger "github.com/kalmecak/go-error-logger"
)

// Delete Elimina una wishlist
func Delete(c *fiber.Ctx) error {

	// extraemos los datos del request
	id := c.Context().UserValue("paramID").(*sql.UID)
	user := c.Context().UserValue("userID").(*sql.UID)

	wl := sql.Wishlist{
		ID: *id,
	}

	db, err := sql.GormDB()
	if err != nil {

		logger.Error(err, "api.withlist.Delete.GormDB")
		var res common.Response
		res.InternalError("could not connect to BD", "")
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}

	// Ejecutamos el soft delete
	tx := db.Where("user_id = ?", user).Delete(&wl)
	if err := tx.Error; err != nil {

		logger.Error(err, "api.withlist.Delete.Where")
		var res common.Response
		res.InternalError("error deleting wishlist", "")
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}

	// Revisamos las filas afectadas, si no se eliminó, lo más seguro es que no existe
	// la wishlist o el usuario no la tiene asiganda
	if tx.RowsAffected == 0 {

		logger.Message("No se elíminó la wishlist", "api.withlist.Delete.RowsAffected")
		var res common.Response
		res.NotFound("wishlist not found")
		return c.Status(fiber.StatusNotFound).JSON(&res)
	}
	res := common.Response{}
	res.Ok("")
	return c.Status(fiber.StatusOK).JSON(&res)
}
