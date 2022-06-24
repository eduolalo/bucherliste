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
		if query := q.GetGoogleQuery(); query != "AIzaSyAqlSYDVik9vOBuLLhpIK_TNv7bh-VbHrk&q=" {

			t.Errorf("El query solo debería contener 'q=', se encontró: %s", query)
		}
	})

	t.Run("GetGoogleQuery_Client_Key", func(t *testing.T) {
		q := Query{
			Key: "AIzaSyAqlSYDVik9vOBuLLhpIK_TNv7bh-EDuaD",
		}
		if query := q.GetGoogleQuery(); query != "AIzaSyAqlSYDVik9vOBuLLhpIK_TNv7bh-EDuaD&q=" {

			t.Errorf("El query solo debería contener 'q=', se encontró: %s", query)
		}
	})
}
