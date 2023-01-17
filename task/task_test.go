package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTask(t *testing.T) {
	task := NewTask("taskName", Image{Name: "imageName", Tag: "0.0.1"})
	want := &Task{
		Name: "taskName",
		Image: Image{
			Name: "imageName",
			Tag:  "0.0.1",
		},
	}

	require.Equal(t, task, want)
}
