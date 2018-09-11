package flix

type User struct {
	ID       int
	Username string
}

type Movie struct {
	ID     int
	Title  string
	UserID int
}

type StorageService interface {
	AddMovie(string) error
	GetMovie(int) (*Movie, error)
	DeleteMovie(int) error
}
