package common

import "os"

// secret es la vairable que almacena la clave secreta para el JWT HMAC
var secret = func() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}()
