package sql

import "github.com/google/uuid"

// UIDFromString Genera un UID a partir de un string
func UIDFromString(s string) (UID, error) {
	id, err := uuid.Parse(s)
	return UID(id), err
}
