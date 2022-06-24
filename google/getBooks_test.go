package google

import (
	"testing"
)

func TestGetBooks(t *testing.T) {

	t.Run("GetBooks", func(t *testing.T) {
		q := Query{
			Author:    "Camilla",
			Title:     "perdonan",
			Publisher: "planeta",
		}
		_, err := GetBooks(q.GetGoogleQuery())
		if err != nil {
			t.Error(err)
		}
	})
}
