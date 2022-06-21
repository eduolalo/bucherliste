package sql

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

// UID customización para trabajar con UUDIDs en MySQL
type UID uuid.UUID

//String Regresa la represtnacion string del UID
func (u UID) String() string {
	return uuid.UUID(u).String()
}

//GormDataType Indica a Gorm el tipo de dato para la BD
func (u UID) GormDataType() string {
	return "binary(16)"
}

// MarshalJSON Regresa la representación para JSON del UID
func (u UID) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(u)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}

// UnmarshalJSON Parsea de JSON a UID
func (u *UID) UnmarshalJSON(by []byte) error {
	s, err := uuid.ParseBytes(by)
	*u = UID(s)
	return err
}

// Scan Implementación del scaner SQL
func (u *UID) Scan(value interface{}) error {

	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*u = UID(parseByte)
	return err
}

// Value Mapea el UID a []byte para alacenar en la BD como []bytes
func (u UID) Value() (driver.Value, error) {
	return uuid.UUID(u).MarshalBinary()
}
