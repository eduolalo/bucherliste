package sql

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Wishlist es la estructura para trabajar con una lista de deseos del usuario
type Wishlist struct {
	gorm.Model

	// Identificador de la lista
	ID uuid.UUID `gorm:"type:BINARY(16);not null;unique;default:(UUID_TO_BIN(UUID()))"`
	// Nombre de la lista
	Name string `gorm:"size:64;not null" validate:"required,max=64" json:"name"`

	/*                                    Relaciones                                    */
	// Wishlists belongs to a User
	UserID uuid.UUID `gorm:"type:BINARY(16);not null"`
	User   User

	// Wishlists has many Books
	Books []Book `gorm:"foreignkey:WishlistID"`
}
