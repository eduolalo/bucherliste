package google

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/kalmecak/bucherliste/common"
	logger "github.com/kalmecak/go-error-logger"
)

// GetBooks Ejecuta la búsqueda de los libros solicitados en el querystring
func GetBooks(q string) ([]Book, error) {

	// construimos el request
	req, err := http.NewRequest("GET", (os.Getenv("API_URL") + q), nil)
	if err != nil {
		return nil, err
	}

	// creamos un context con timeout de 10 segundos, (Google responde en menos de 1)
	ctx, cancel := context.WithTimeout(req.Context(), 10*time.Second)
	defer cancel()
	req = req.WithContext(ctx)

	// Creamos el cliente http y ejecutamos el request
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// leemos el body del response
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// revisamos la respuesta de Google
	if res.StatusCode != 200 {

		logger.Message(string(b), "api.google.GetBooks.http.Do")
		return nil, errors.New("google api error, check your parameters")
	}
	// parseamos el body del response
	body := booksBody{}
	if err := common.JSON.Unmarshal(b, &body); err != nil {
		return nil, err
	}

	var books []Book
	for _, b := range body.Items {

		books = append(books, b.ToBook())
	}

	return books, nil
}

// body estructura temporal para parsear el body de la respuesta de google
type booksBody struct {
	Items []GBook `json:"items"`
}
