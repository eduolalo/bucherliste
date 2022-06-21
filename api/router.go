package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/api/user"
	"github.com/kalmecak/bucherliste/api/validate"
	"github.com/kalmecak/bucherliste/api/wishlist"
)

// Router maneja las peticiones y agrega los middlewares para cada caso
func Router(app *fiber.App) {

	/**************************************************************************/
	/*                                Usuarios                                */
	/**************************************************************************/
	// Registro de usuario
	app.Post("/signup", validate.UserBody, user.SignUp)
	// Sesión de usuario
	app.Post("/login", validate.UserBody, user.LogIn)
	/**************************************************************************/
	/*                                Wishlists                               */
	/**************************************************************************/
	// Crear wishlist
	app.Post("/wishlist", validate.IsLogged, validate.CreateBody, wishlist.Post)
	// Obtener wishlist
	app.Get("/wishlists", validate.IsLogged, wishlist.Get)
	// Obtener contenido de wishlist
	app.Get("/wishlist/:id", validate.IsLogged, validate.IDParam, wishlist.Wishlist)
	// Actualizar wishlist
	app.Put("/wishlist/:id", validate.IsLogged, validate.IDParam, wishlist.Put)
	// Eliminar wishlist
	app.Delete("/wishlist/:id", validate.IsLogged, validate.IDParam, wishlist.Delete)

	/**************************************************************************/
	/*                                 Libros                                 */
	/**************************************************************************/
	// Búsqueda de libros
	app.Get("/books", wishlist.Wishlist)
}
