package postgres

import (
	"database/sql"
	"log"

	"github.com/drewvanstone/flix"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewDB() *DB {
	db, err := sql.Open("postgres", "postgres://postgres:@localhost/flix_development?sslmode=disable")
	if err != nil {
		log.Panic(err)

	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	return &DB{db}
}

func (db DB) AddMovie(desc string) error {
	stmt, err := db.Prepare("INSERT INTO movies(title) VALUES($1)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(desc)
	if err != nil {
		return err
	}
	return nil
}

func (db DB) GetMovie(id int) (*flix.Movie, error) {
	m := flix.Movie{}
	row := db.QueryRow("SELECT * FROM movies WHERE id = $1", id)
	if err := row.Scan(&m.ID, &m.Title, &m.UserID); err != nil {
		return nil, err
	}
	return &m, nil
}

func (db DB) DeleteMovie(id int) error {
	stmt, err := db.Prepare("DELETE FROM movies WHERE id = $1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
