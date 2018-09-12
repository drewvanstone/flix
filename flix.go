package flix

type User struct {
	ID       int
	Username string
}

type Movie struct {
	ID    int
	Title string
}

type StorageService interface {
	AddMovie(string) error
	GetMovie(int) (*Movie, error)
	GetMovies() ([]Movie, error)
	DeleteMovie(int) error
}
