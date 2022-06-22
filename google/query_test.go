package google

import "testing"

// TestQuery set de testing para la estructura Query
func TestQuery(t *testing.T) {

	t.Run("GetGoogleQuery_All", func(t *testing.T) {
		q := Query{
			Author:    "Camilla",
			Title:     "perdonan",
			Publisher: "planeta",
		}
		if q.GetGoogleQuery() == "" {
			t.Error("El query no debe ser vacio")
		}
	})

	t.Run("GetGoogleQuery_None", func(t *testing.T) {
		q := Query{}
		if q.GetGoogleQuery() != "q=" {
			t.Error("El query solo debería contener 'q='")
		}
	})
}
