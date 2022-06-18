package common

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	logger "github.com/kalmecak/go-error-logger"
	"github.com/pascaldekloe/jwt"
)

// Payload representación del payload del jwt de authorización
type Payload struct {
	ISS string `json:"iss" validate:"required,numeric"`
	Nbf int    `json:"nbf" validate:"required,numeric"`
	Iat int    `json:"iat" validate:"required,numeric"`
	Exp int    `json:"exp" validate:"required,numeric"`
	Ref string `json:"ref" validate:"required,uuid"`
}

// Unmarshall usa la sección de payload del jwt para integrar la información a la estructura
// Nota: ESTE MÉTODO DEBE USARSE DESPUÉS DE HABER VALIDADO EL JWT
func (s *Payload) Unmarshall(body64 string) error {

	if l := len(body64) % 4; l > 0 {
		body64 += strings.Repeat("=", 4-l)
	}
	bodyByte, err := base64.URLEncoding.DecodeString(body64)
	if err != nil {

		logger.Error(err, "common.payload.Unmarshall.DecodeString")
		return errors.New("payload is not base64")
	}

	if err = json.Unmarshal(bodyByte, s); err != nil {

		logger.Error(err, "structs.Payload.Unmarshall.json.Unmarshall")
		return errors.New("payloadmalformed")
	}
	return nil
}

// IsValid revisa si todos los tags se cumplen
func (s Payload) IsValid() (err error) {

	validate := validator.New()
	err = validate.Struct(s)
	if err != nil {

		errs := err.(validator.ValidationErrors)
		// Vamos a iterar sobre el arreglo de errores
		// para obtener el string de cada uno y mandarlos
		// en un solo string
		var errString strings.Builder
		size := len(errs) - 1
		for i, e := range errs {

			errString.WriteString(e.Error())
			if i < size {

				errString.WriteString("; ")
			}
		}
		err = errors.New(errString.String())
	}
	return
}

// Validate Revisa la validéz del JWT y en caso de no tener errores, integra la
// información al payload
func (s *Payload) Validate(t string) error {

	// Validación del token
	claims, err := jwt.HMACCheck([]byte(t), secret)
	if err != nil {

		logger.Error(err, "common.payload.Validate.HMACCheck")
		return err
	}

	// Validación del tiempo del token
	if !claims.Valid(time.Now()) {

		err := errors.New("token expired")
		logger.Error(err, "common.payload.Validate.HMACCheck")
		return err
	}

	// Se separa el payload del token y se intenta integrar la información
	// en caso de error, se regresará el del método unmarshall
	sliced := strings.Split(t, ".")
	return s.Unmarshall(sliced[1])
}
