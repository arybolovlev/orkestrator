package job

import (
	"testing"

	"github.com/arybolovlev/orkestrator/task"

	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
	job := NewJob("jobName", []task.Task{{Name: "taskName", Image: task.Image{Name: "imageName", Tag: "0.0.1"}}})
	want := &Job{
		Name: "jobName",
		Tasks: []task.Task{
			{
				Name: "taskName",
				Image: task.Image{
					Name: "imageName",
					Tag:  "0.0.1",
				},
			},
		},
	}

	require.Equal(t, job, want)
}
