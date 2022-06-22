package sql

import (
	"crypto/sha256"

	"github.com/kalmecak/bucherliste/common"
	"gorm.io/gorm"
)

// User es la estructura para trabajar con los datos de un usuario
type User struct {
	gorm.Model

	// Identificador público del usuario
	ID UID `gorm:"primaryKey;default:(UUID_TO_BIN(UUID()))"`
	// Nombre(s) del usuario
	Username string `gorm:"size:64;UNIQUE" validate:"required,max=64" json:"username"`
	// Password
	Password string `gorm:"-" validate:"required,max=64"  json:"password"`
	Hash     []byte `gorm:"type:BINARY(32);not null" json:"-"`

	/*                                    Relaciones                                    */
	// Users has many Wishlists
	Wishlists []Wishlist `gorm:"foreignkey:UserID"`
}

/**************************************************************************/
/*                                Métodos                                 */
/**************************************************************************/

// Unmarshal acomoda los datos del body en la estructura
func (u *User) Unmarshal(body []byte) error {

	if err := common.JSON.Unmarshal(body, u); err != nil {

		return err
	}
	return nil
}

// Validate corre las validaciones de la estructura
func (u *User) Validate() error {
	return common.ValidateStruct(u, "User.")
}

// BuilHash construye el hash del password para almacenar en la base de datos
func (u *User) BuilHash() error {

	raw := u.Username + "@" + u.Password
	hash := sha256.New()
	if _, err := hash.Write([]byte(raw)); err != nil {
		return err
	}
	resume := hash.Sum(nil)
	u.Hash = resume
	return nil
}

/**************************************************************************/
/*                                DB Hooks                                */
/**************************************************************************/

// BeforeCreate ejecuta lógica antes de crear un nuevo usuario
func (u *User) BeforeCreate(tx *gorm.DB) error {

	return u.BuilHash()
}
