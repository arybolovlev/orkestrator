package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTask(t *testing.T) {
	task := NewTask("name", Image{Name: "self", Tag: "0.0.1"})
	want := &Task{
		Name: "name",
		Image: Image{
			Name: "self",
			Tag:  "0.0.1",
		},
	}

	require.Equal(t, task, want)
}
