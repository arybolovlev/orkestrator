package task

import (
	"testing"

	"github.com/arybolovlev/orkestrator/helpers"
	"github.com/stretchr/testify/require"
)

func TestNewTask(t *testing.T) {
	task := NewTask("taskName", Image{Name: "imageName", Tag: helpers.PointerOf("0.0.1")})
	want := &Task{
		Name: "taskName",
		Image: Image{
			Name: "imageName",
			Tag:  helpers.PointerOf("0.0.1"),
		},
	}

	require.Equal(t, task, want)
}

func TestValidateTaskEmptyTaskName(t *testing.T) {
	task := &Task{
		Name: "",
		Image: Image{
			Name: "imageName",
			Tag:  helpers.PointerOf("0.0.1"),
		},
	}
	require.Error(t, task.Validate())
}

func TestValidateTaskEmptyImageName(t *testing.T) {
	task := &Task{
		Name: "taskName",
		Image: Image{
			Name: "",
			Tag:  helpers.PointerOf("0.0.1"),
		},
	}
	require.Error(t, task.Validate())
}

func TestValidateTaskEmptyImageTag(t *testing.T) {
	task := &Task{
		Name: "taskName",
		Image: Image{
			Name: "imageName",
		},
	}
	require.NoError(t, task.Validate())
}

func TestValidateTask(t *testing.T) {
	task := &Task{
		Name: "taskName",
		Image: Image{
			Name: "imageName",
			Tag:  helpers.PointerOf("0.0.1"),
		},
	}
	require.NoError(t, task.Validate())
}
