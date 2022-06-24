package google

import (
	"reflect"
	"strings"
)

// Query es la estructura para trabajar con la query del request para búsqueda de libros
type Query struct {
	Key       string `query:"key" search:"" validate:"omitempty,max=64"`
	Author    string `query:"a" search:"+inauthor:" validate:"omitempty,max=64"`
	Title     string `query:"t" search:"+intitle:" validate:"omitempty,max=64"`
	Publisher string `query:"p" search:"+inpublisher:" validate:"omitempty,max=64"`
}

/**************************************************************************/
/*                                Métodos                                 */
/**************************************************************************/

// GetGoogleQuery genera el query string para la búsqueda en Google Books
func (q *Query) GetGoogleQuery() string {

	// generamos el builder para la query param
	var qprm strings.Builder

	// Si en el query hay una key, la agregamos al query param
	if q.Key != "" {

		qprm.WriteString(q.Key)
	} else {

		// Si no hay key, usamos la propia
		qprm.WriteString(apiKey)
	}
	qprm.WriteString("&")
	qprm.WriteString("q=")

	// iteramos por los campos de la estructura
	fields := reflect.TypeOf(*q)
	values := reflect.ValueOf(*q)
	l := fields.NumField()
	// empezamos del index 1 porque el campo Key no se incluye en la query dinámica
	for i := 1; i < l; i++ {

		// si el valor del campo no es vacio, lo agregamos al qprm
		if v := values.Field(i).String(); v != "" {

			// obtenemos el tag del campo para buscar en google
			qprm.WriteString(fields.Field(i).Tag.Get("search"))
			// obtenemos el valor para buscar en google
			qprm.WriteString(v)
		}
	}
	return qprm.String()
}
