package sql

import (
	"github.com/kalmecak/bucherliste/common"
	"gorm.io/gorm"
)

// Wishlist es la estructura para trabajar con una lista de deseos del usuario
type Wishlist struct {
	gorm.Model

	// Identificador de la lista
	ID UID `gorm:"primaryKey;default:(UUID_TO_BIN(UUID()))"`
	// Nombre de la lista
	Name string `gorm:"size:64;not null" validate:"required,max=64" json:"name"`

	/*                                    Relaciones                                    */
	// Wishlists belongs to a User
	UserID UID  `gorm:"not null" json:"-" `
	User   User `json:"-" validate:"-"`

	// Wishlists has many Books
	Books []Book `gorm:"foreignkey:WishlistID" validate:"-"`
}

/**************************************************************************/
/*                                Métodos                                 */
/**************************************************************************/

// Unmarshal acomoda los datos del body en la estructura
func (w *Wishlist) Unmarshal(body []byte) error {

	if err := common.JSON.Unmarshal(body, w); err != nil {

		return err
	}
	return nil
}

// Validate corre las validaciones de la estructura
func (w *Wishlist) Validate() error {
	return common.ValidateStruct(w, "Wishlist.")
}
