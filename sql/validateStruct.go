package sql

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

// validateStruct corre validaciones de las estructuras usando sus tags
func validateStruct(s interface{}, name string) error {

	validate := validator.New()
	err := validate.Struct(s)
	if err == nil {

		return nil
	}

	errs := err.(validator.ValidationErrors)
	var errString strings.Builder
	size := len(errs) - 1

	for i, e := range errs {

		errString.WriteString(e.Error())
		if i < size {

			errString.WriteString("; ")
		}
	}
	str := errString.String()
	str = strings.ReplaceAll(str, name, "")
	str = strings.ToLower(str)
	return errors.New(str)
}
