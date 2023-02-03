package task

import (
	"errors"
)

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

func (t *Task) Validate() error {
	if t.Name == "" {
		return errors.New("Task Name must be set")
	}

	if t.Image.Name == "" {
		return errors.New("Image Name must be set")
	}

	return nil
}
