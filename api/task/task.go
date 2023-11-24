package task

import (
	"errors"
)

type Image struct {
	Name string  `hcl:"name"`
	Tag  *string `hcl:"tag"`
}

type Task struct {
	Name  string `hcl:"name,label"`
	Image Image  `hcl:"image,block"`
}

func NewTask(n string, i Image) *Task {
	if i.Tag == nil {
		t := "latest"
		i.Tag = &t
	}

	return &Task{
		Name:  n,
		Image: i,
	}
}

func (t *Task) Validate() error {
	if t.Name == "" {
		return errors.New("Task Name must be set")
	}

	return t.Image.Validate()
}

func (i *Image) Validate() error {
	if i.Name == "" {
		return errors.New("Image Name must be set")
	}

	return nil
}
