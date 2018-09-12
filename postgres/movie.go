package postgres

import "github.com/drewvanstone/flix"

func (db DB) AddMovie(title string) error {
	stmt, err := db.Prepare("INSERT INTO movies(title) VALUES($1)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(title)
	if err != nil {
		return err
	}
	return nil
}

func (db DB) GetMovie(id int) (*flix.Movie, error) {
	m := flix.Movie{}
	row := db.QueryRow("SELECT * FROM movies WHERE id = $1", id)
	if err := row.Scan(&m.ID, &m.Title); err != nil {
		return nil, err
	}
	return &m, nil
}

func (db DB) GetMovies() ([]flix.Movie, error) {
	var movies []flix.Movie

	rows, err := db.Query("SELECT * FROM movies")
	if err != nil {
		return movies, err
	}
	defer rows.Close()
	for rows.Next() {
		m := flix.Movie{}
		err := rows.Scan(&m.ID, &m.Title)
		if err != nil {
			return movies, err
		}
		movies = append(movies, m)
	}
	err = rows.Err()
	if err != nil {
		return movies, err
	}

	return movies, nil
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
