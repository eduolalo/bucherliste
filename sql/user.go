package sql

import "gorm.io/gorm"

// User es la estructura para trabajar con los datos de un usuario
type User struct {
	gorm.Model

	// Identificador público del usuario
	ID []byte `gorm:"type:BINARY(16);not null;unique;default:(UUID_TO_BIN(UUID()))"`
	// Nombre(s) del usuario
	Username string `gorm:"size:64;not null" validate:"required,max=64" json:"username"`
	// Password
	Password string `gorm:"-" validate:"required,max=64"  json:"password"`
	Hash     []byte `gorm:"type:BINARY(32);not null" json:"-"`

	/*                                    Relaciones                                    */
	// Users has many Wishlists
	WishList []WishList `gorm:"foreignkey:UserID"`
}
