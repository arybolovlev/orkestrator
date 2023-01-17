package job

import (
	"github.com/arybolovlev/orkestrator/task"
)

type Job struct {
	Name  string
	Tasks []task.Task
}

func NewJob(n string, t []task.Task) *Job {
	return &Job{
		Name:  n,
		Tasks: t,
	}
}
