package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/api/user"
	"github.com/kalmecak/bucherliste/api/wishlist"
)

// Router maneja las peticiones y agrega los middlewares para cada caso
func Router(app *fiber.App) {

	/**************************************************************************/
	/*                                Usuarios                                */
	/**************************************************************************/
	// Registro de usuario
	app.Post("/signup", user.SignUp)
	// Sesión de usuario
	app.Post("/login", user.LogIn)
	/**************************************************************************/
	/*                                Wishlists                               */
	/**************************************************************************/
	// Crear wishlist
	app.Post("/wishlist", wishlist.Post)
	// Obtener wishlist
	app.Get("/wishlists", wishlist.Get)
	// Obtener contenido de wishlist
	app.Get("/wishlist/:id", wishlist.Books)
	// Actualizar wishlist
	app.Put("/wishlist/:id", wishlist.Put)
	// Eliminar wishlist
	app.Delete("/wishlist/:id", wishlist.Delete)

	/**************************************************************************/
	/*                                 Libros                                 */
	/**************************************************************************/
	// Búsqueda de libros
	app.Get("/books", wishlist.Books)
}
