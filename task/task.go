package task

type Image struct {
	Name string
	Tag  string
}

type Task struct {
	Name  string
	Image Image
}

func NewTask(n string, i Image) *Task {
	return &Task{
		Name:  n,
		Image: i,
	}
}
