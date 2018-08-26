package minerva

type User struct {
	ID       int
	Username string
	TaskID   int
}

type Task struct {
	ID          int
	Description string
	Status      string
}

type StorageService interface {
	Create(string) error
	Read(int) (*Task, error)
	Delete(int) error
}
