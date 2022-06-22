package google

import "os"

// apiKey expone la llave de API de Google Books
var apiKey = func() string {

	return os.Getenv("API_KEY")
}()
