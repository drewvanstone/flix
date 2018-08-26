package postgres

import (
	"database/sql"
	"log"

	"github.com/drewvanstone/minerva"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewDB() *DB {
	db, err := sql.Open("postgres", "postgres://postgres:@localhost/minerva_development?sslmode=disable")
	if err != nil {
		log.Panic(err)

	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	return &DB{db}
}

func (db DB) Create(desc string) error {
	stmt, err := db.Prepare("INSERT INTO tasks(description) VALUES($1)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(desc)
	if err != nil {
		return err
	}
	return nil
}

func (db DB) Read(id int) (*minerva.Task, error) {
	t := minerva.Task{}
	row := db.QueryRow("SELECT * FROM tasks WHERE id = $1", id)
	if err := row.Scan(&t.ID, &t.Description, &t.Status); err != nil {
		return nil, err
	}
	return &t, nil
}

func (db DB) Delete(id int) error {
	stmt, err := db.Prepare("DELETE FROM tasks WHERE id = $1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
