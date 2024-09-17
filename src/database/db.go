package database

import (
	"database/sql"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
	err  error
)

func ConectarDB() *sql.DB {
	once.Do(func() {
		if db, err = sql.Open("postgres", os.Getenv("stringDB")); err != nil {
			log.Fatal()
		}

		if err = db.Ping(); err != nil {
			log.Fatal(err.Error())
		}
	})
	return db
}
