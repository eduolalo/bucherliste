package sql

import (
	"gorm.io/gorm"
)

// Book es la estructura para trabajar con los datos de un libro
type Book struct {
	gorm.Model

	// Identificador del libro
	ID UID `gorm:"primaryKey;default:(UUID_TO_BIN(UUID()))"`
	// Identificador en google books
	GID string `gorm:"size:64;UNIQUE" validate:"required,max=64" json:"gid"`
	// Titulo del libro
	Title string `gorm:"size:64;not null" validate:"required,max=128" json:"title"`
	// Autor del libro
	Authors string `gorm:"size:256;not null" validate:"required,max=256" json:"authors"`
	// Publicador del libro
	Publisher string `gorm:"size:128;not null" validate:"required,max=128" json:"publisher"`

	/*                                    Relaciones                                    */
	// Book pertenece a una wishlist
	WishlistID UID `json:"-"`
}
