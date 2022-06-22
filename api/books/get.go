package books

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/bucherliste/common"
	"github.com/kalmecak/bucherliste/google"
	logger "github.com/kalmecak/go-error-logger"
)

// Get obtiene una lista de libros según los parámetros de búsqueda
func Get(c *fiber.Ctx) error {

	q := google.Query{}
	if err := c.QueryParser(&q); err != nil {

		logger.Error(err, "api.books.Get.queryParser")
		res := booksRes{}
		res.Ok("")
		return c.Status(fiber.StatusOK).JSON(&res)
	}

	books, err := google.GetBooks(q.GetGoogleQuery())
	if err != nil {

		logger.Error(err, "api.books.Get.getBooks")
		res := booksRes{}
		res.InternalError("could not get books", "")
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}
	res := booksRes{
		Books: books,
	}

	res.Ok("")
	return c.Status(fiber.StatusOK).JSON(&res)
}

// booksRes es la estructura de respuasta para la búsqueda de libros
type booksRes struct {
	common.Response
	Books []google.Book `json:"books"`
}
