package validate

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
	"github.com/kalmecak/bucherliste/sql"
	logger "github.com/kalmecak/go-error-logger"
)

// ValidateWishlistBody valida que el request tenga los da
// datos necesarios para modificar la lista de deseos
func ValidateWishlistBody(c *fiber.Ctx) error {

	// Extraemos el body del request
	var body wishlistBody
	if err := common.JSON.Unmarshal(c.Body(), &body); err != nil {

		logger.Error(err, "api.validate.ValidateWishlistBody.unmarshal")
		var res common.Response
		res.InternalError("error parsing body", "")
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}

	// Ejecutamos las validacions de los objetos books
	if err := common.ValidateStruct(body, "wishlistBody"); err != nil {

		logger.Error(err, "api.validate.ValidateWishlistBody.ValidateStruct")
		var res common.Response
		res.BadRequest("body is not valid")
		return c.Status(fiber.StatusBadRequest).JSON(&res)
	}

	// Guardamos el body en el contexto
	remove := (body.Delete == "true")
	c.Context().SetUserValue("books", &body.Books)
	c.Context().SetUserValue("del", &remove)
	return c.Next()
}

// wishlistBody es la estructura para trabajar el request de modificación de la lista
// de deseos
type wishlistBody struct {
	Delete string     `json:"delete" validate:"oneof=true false"`
	Books  []sql.Book `json:"books" validate:"required"`
}
