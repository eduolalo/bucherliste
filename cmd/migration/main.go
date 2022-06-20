package main

import (
	"log"

	"github.com/kalmecak/bucherliste/sql"
)

func main() {

	db, err := sql.GormDB()
	if err != nil {
		log.Panic(err)
		return
	}

	err = db.AutoMigrate(
		&sql.User{},
		&sql.WishList{},
		&sql.Book{},
	)
	if err != nil {
		log.Panic(err)
	}
}
