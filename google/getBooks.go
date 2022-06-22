package google

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/kalmecak/bucherliste/common"
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
