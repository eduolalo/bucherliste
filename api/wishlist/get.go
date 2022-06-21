package wishlist

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
	"github.com/kalmecak/bucherliste/sql"
	logger "github.com/kalmecak/go-error-logger"
	"gorm.io/gorm"
)

// Get Obtiene todas las wishlists del usuario
func Get(c *fiber.Ctx) error {

	userID := c.Context().UserValue("userID").(*sql.UID)
	user := sql.User{
		ID: *userID,
	}

	// Conectamos con la BD
	db, err := sql.GormDB()
	if err != nil {

		logger.Error(err, "api.wishlist.Get.db.GormDB")
		var res common.Response
		res.InternalError("could not connect to BD", "")
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}
	var wls []sql.Wishlist
	err = db.Model(&user).Association("Wishlists").Find(&wls)
	if err != nil {

		res := listRes{
			Wishlists: []sql.Wishlist{},
		}
		// Handleamos el error de not found
		if err == gorm.ErrRecordNotFound {
			res.NotFound("user have not any wishlist")
			c.Status(fiber.StatusNotFound)
		} else {

			// manejo de cualquier otro error
			logger.Error(err, "api.wishlist.Get.db.Association.Find")
			res.InternalError("could not get wishlists", "")
			c.Status(fiber.StatusInternalServerError)
		}
		return c.JSON(&res)
	}
	res := listRes{
		Wishlists: wls,
	}
	res.Ok("")
	return c.Status(fiber.StatusOK).JSON(&res)
}

// listRes Respuesta para la peticion de lista de deseos del usuario
type listRes struct {
	common.Response
	Wishlists []sql.Wishlist `json:"wishlists"`
}
