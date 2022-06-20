package sql

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// GormDB genera la conexión la base de datos
func GormDB() (*gorm.DB, error) {

	// Si ya existe una instancia de la base de datos, se devuelve
	if GDB != nil {

		return GDB, nil
	}

	/*                                  Nueva conexión                                  */

	// Se genera una nueva conexión
	str := os.Getenv("DB_STRING")
	// COnfiguaración del logger de GORM
	conf := logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
		Colorful:                  false,       // Disable color
	}

	// Setteamos el nivel de error dependiendo de las variables de entorno
	// en productivo deve ser "Error"
	switch os.Getenv("LOG_LEVEL") {
	case "Error":
		conf.LogLevel = logger.Error
	case "Warn":
		conf.LogLevel = logger.Warn
	default:
		conf.LogLevel = logger.Info

	}
	// re registra el logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		conf,
	)
	// Se crea la conexión a la base de datos
	db, err := gorm.Open(mysql.Open(str), &gorm.Config{
		Logger: newLogger,
	})
	DB = db
	return db, err
}
