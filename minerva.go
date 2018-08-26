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
	CreateTask(string) error
	ReadTask(int) (*Task, error)
	DeleteTask(int) error
}
