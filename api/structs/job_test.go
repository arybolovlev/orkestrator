package structs

import (
	"testing"

	"github.com/arybolovlev/orkestrator/api/task"
	"github.com/arybolovlev/orkestrator/helpers"

	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
	job := NewJob("jobName", []task.Task{{Name: "taskName", Image: task.Image{Name: "imageName", Tag: helpers.PointerOf("0.0.1")}}})
	want := &Job{
		Name: "jobName",
		Tasks: []task.Task{
			{
				Name: "taskName",
				Image: task.Image{
					Name: "imageName",
					Tag:  helpers.PointerOf("0.0.1"),
				},
			},
		},
	}

	require.Equal(t, job, want)
}

func TestValidateJobEmptyJobName(t *testing.T) {
	job := &Job{
		Name: "",
		Tasks: []task.Task{
			{
				Name: "taskName",
				Image: task.Image{
					Name: "imageName",
					Tag:  helpers.PointerOf("0.0.1"),
				},
			},
		},
	}
	require.Error(t, job.Validate())
}

func TestValidateJobNoTasks(t *testing.T) {
	job := &Job{
		Name:  "jobName",
		Tasks: []task.Task{},
	}
	require.Error(t, job.Validate())
}

func TestValidateJobValidateTaskEmptyTaskName(t *testing.T) {
	job := &Job{
		Name: "",
		Tasks: []task.Task{
			{
				Name: "",
				Image: task.Image{
					Name: "imageName",
					Tag:  helpers.PointerOf("0.0.1"),
				},
			},
		},
	}
	require.Error(t, job.Validate())
}

func TestValidateJobValidateTaskEmptyImageName(t *testing.T) {
	job := &Job{
		Name: "",
		Tasks: []task.Task{
			{
				Name: "taskName",
				Image: task.Image{
					Name: "",
					Tag:  helpers.PointerOf("0.0.1"),
				},
			},
		},
	}
	require.Error(t, job.Validate())
}
