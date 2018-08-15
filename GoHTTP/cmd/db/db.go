package db

import (
	"github.com/go-pg/pg"
	"github.com/labstack/gommon/log"
)



func Connect() *pg.DB{
	opts := &pg.Options{
		User:	"postgres",
		Password:	"cvbmnb",
		Database:	"library",
	}

	var db = pg.Connect(opts)

	if db == nil{
		log.Printf("Failed to connect to database!")
		return nil
	}
	log.Printf("Connected to database!")
	return db
}

