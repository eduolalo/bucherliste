package migration

import (
	"time"

	"github.com/kalmecak/bucherliste/sql"
)

// Start ejecuta la creación de la base de datos.
func Start() error {

	// Aunque el contenedor de la BD inicia promero, tarda unos segundo en iniciar el
	// servicio, por lo cual ejecutamos la migración con un tiempo de espera.
	time.Sleep(6 * time.Second)

	// ejecutamos el migrate
	db, err := sql.GormDB()
	if err != nil {
		return err
	}

	return db.AutoMigrate(
		&sql.User{},
		&sql.Wishlist{},
		&sql.Book{},
	)
}
