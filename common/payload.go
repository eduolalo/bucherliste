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
	ISS string  `json:"iss" validate:"required,hostname"`
	Nbf float64 `json:"nbf" validate:"required,numeric"`
	Iat float64 `json:"iat" validate:"required,numeric"`
	Exp float64 `json:"exp" validate:"required,numeric"`
	Ref string  `json:"ref" validate:"required,uuid"`
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

// JWT Genera un JWT de sesión con la información del payload
func (s *Payload) JWT() (string, error) {

	var c jwt.Claims
	// seteamos los claims
	c.Issuer = "bucherlist.com"
	c.Set = map[string]interface{}{
		"ref": s.Ref,
	}
	now := time.Now()
	from := jwt.NewNumericTime(now)
	c.NotBefore = from
	c.Issued = from
	// Le damos 5 minutos de expiración
	c.Expires = jwt.NewNumericTime(now.Add(time.Minute * time.Duration(jwtTTL)))

	t, err := c.HMACSign(jwt.HS256, secret)
	if err != nil {

		logger.Error(err, "common.payload.JWT.HMACSign")
		return "", err
	}
	return string(t), nil
}
