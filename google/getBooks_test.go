package google

import (
	"log"
	"testing"
)

func TestGetBooks(t *testing.T) {

	t.Run("GetBooks", func(t *testing.T) {
		q := Query{
			Author:    "Camilla",
			Title:     "perdonan",
			Publisher: "planeta",
		}
		s, err := GetBooks(q.GetGoogleQuery())
		if err != nil {
			t.Error(err)
		} else {
			log.Printf("búsqueda: %+v", s)
		}
	})
}
