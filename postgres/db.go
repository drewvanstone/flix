package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewDB() *DB {
	db, err := sql.Open("postgres", "postgres://postgres:@localhost/flix_development?sslmode=disable")
	if err != nil {
		log.Fatalf("Could not connect to database. Got error %v\n", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Could not ping database. Got error %v\n", err)
	}
	return &DB{db}
}
