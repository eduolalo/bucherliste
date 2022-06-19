package db

import "gorm.io/gorm"

// WishList es la estructura para trabajar con una lista de deseos del usuario
type WishList struct {
	gorm.Model

	// Identificador de la lista
	ID []byte `gorm:"type:BINARY(16);not null;unique;default:(UUID_TO_BIN(UUID()))"`
	// Nombre de la lista
	Name string `gorm:"size:64;not null" validate:"required,max=64" json:"name"`

	/*                                    Relaciones                                    */
	// Wishlists belongs to a User
	UserID []byte `gorm:"type:BINARY(16);not null"`
	User   User

	// Wishlists has many Books
	Books []Book `gorm:"foreignkey:WishListID"`
}
