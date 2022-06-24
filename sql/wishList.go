package sql

import (
	"log"

	"github.com/kalmecak/bucherliste/common"
	logger "github.com/kalmecak/go-error-logger"
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
	Books []Book `json:"-" validate:"-"`
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

// Append agrega los libros a la lista de deseos
func (w *Wishlist) Append(books *[]Book) error {

	// conectamos a la base de datos
	db, err := GormDB()
	if err != nil {

		logger.Error(err, "sql.wishlist.Append.db.GormDB")
		return err
	}
	return db.Session(&gorm.Session{FullSaveAssociations: true}).
		Model(&w).Association("Books").Append(*books)
}

// Delete recibe un arreglo de books y los elimina de la lista de deseos
func (w *Wishlist) Delete(books *[]Book) error {

	// conectamos a la base de datos
	db, err := GormDB()
	if err != nil {

		return err
	}

	gids := []string{}
	for _, book := range *books {

		gids = append(gids, book.GID)
	}

	tx := db.Where("g_id in ?", gids).
		Where("wishlist_id = ?", w.ID).Find(books)
	if err := tx.Error; err != nil {

		return err
	}

	tx = db.Delete(books)
    if err := tx.Error; err != nil {

            return err
    }
	log.Printf("%+v", books)
	return nil

}
